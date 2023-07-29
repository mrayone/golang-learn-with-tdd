package querybuilder

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	ErrorCreateInOperator   error = errors.New("can't cast arry or slice to in operator")
	ErrorInterateInOperator error = errors.New("can't interate an empty array")
	ErrorInvalidOperator    error = errors.New("invalid operator given")
)

type Filter struct {
	Column string
	Op     FilterOperator
	Value  interface{}
}

type FilterOperator string

const (
	None             FilterOperator = ""
	Equal            FilterOperator = "="
	GreaterThan      FilterOperator = ">"
	GreaterEqualThan FilterOperator = ">="
	LessEqualThan    FilterOperator = "<="
	LessThan         FilterOperator = "<"
	In               FilterOperator = "IN"
	IsNull           FilterOperator = "IS NULL"
	IsNotNull        FilterOperator = "IS NOT NULL"
)

type FilterSplitOperator string

const (
	And FilterSplitOperator = "AND"
	Or  FilterSplitOperator = "OR"
)

type QueryBuilder struct {
	sqlRaw     string
	paramCount int
	groupBy    string
	Args       []interface{}
}

func New(sqlBase string) QueryBuilder {
	return QueryBuilder{
		sqlRaw:     sqlBase,
		paramCount: 0,
		Args:       make([]interface{}, 0),
	}
}

// AddFilters with given params
// Add where clause to query
// Available AND boolean operator and comparable FIlterOperators
func (qb *QueryBuilder) AddFilters(filters []Filter, separator FilterSplitOperator) error {
	for _, filter := range filters {
		err := qb.addFilter(filter, separator)
		if err != nil {
			return err
		}
	}
	return nil
}

// AddGoupBy to query
func (qb *QueryBuilder) AddGroupBy(column string) {
	if qb.groupBy != "" {
		qb.groupBy = fmt.Sprintf("%s,%s", qb.groupBy, column)
		return
	}
	qb.groupBy = fmt.Sprintf(`GROUP BY %s`, column)
}

// RawQuery returns raw of query base
func (qb *QueryBuilder) RawQuery() string {
	if qb.groupBy == "" {
		return qb.sqlRaw
	}

	return fmt.Sprintf("%s \n%s", qb.sqlRaw, qb.groupBy)
}

func (qb *QueryBuilder) getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func (qb *QueryBuilder) addFilter(filter Filter, separator FilterSplitOperator) error {
	whereFilters, err := qb.createFilter(filter)
	if err != nil {
		return err
	}

	if qb.paramCount == 1 {
		qb.sqlRaw = fmt.Sprintf("%s \n%s", qb.sqlRaw, "WHERE")
	}

	if qb.paramCount > 1 {
		qb.sqlRaw = fmt.Sprintf("%s \n%s", qb.sqlRaw, separator)
	}

	qb.sqlRaw = fmt.Sprintf("%s %s", qb.sqlRaw, whereFilters)
	return nil
}

func (qb *QueryBuilder) createFilter(filter Filter) (string, error) {
	switch filter.Op {
	case In:
		return qb.inOperator(filter)
	case IsNull, IsNotNull:
		return fmt.Sprintf("%s %s", filter.Column, filter.Op), nil
	case Equal, LessEqualThan, LessThan, GreaterEqualThan, GreaterThan:
		qb.paramCount++
		qb.Args = append(qb.Args, filter.Value)
		return fmt.Sprintf("%s %s @p%d", filter.Column, filter.Op, qb.paramCount), nil
	}

	return "", ErrorInvalidOperator
}

func (qb *QueryBuilder) inOperator(filter Filter) (string, error) {
	value := qb.getValue(filter.Value)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return "", ErrorCreateInOperator
	}

	size := value.Len()
	if size == 0 {
		return "", ErrorCreateInOperator
	}

	modelsParams := make([]string, size)
	for i := 0; i < size; i++ {
		qb.paramCount++
		modelsParams[i] = fmt.Sprint("@p", qb.paramCount)
		if value.Index(i).Kind() == reflect.String {
			qb.Args = append(qb.Args, value.Index(i).String())
		}

		if value.Index(i).Kind() == reflect.Int64 || value.Index(i).Kind() == reflect.Int {
			qb.Args = append(qb.Args, value.Index(i).Int())
		}
		//TODO: another types
	}

	return fmt.Sprintf(`%s %s (%s)`, filter.Column, filter.Op, strings.Join(modelsParams, ",")), nil
}
