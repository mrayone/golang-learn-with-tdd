package querybuilder_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/mrayone/learn-go/querybuilder"
)

const (
	mainQuery string = `
SELECT
	ISNULL(SUM(p.Total), 0),
	ISNULL(COUNT(1), 0)
FROM
	Order p (nolock)
JOIN
	Customer c (nolock) ON c.Id = p.CustomerId`
)

func TestAddFilters(t *testing.T) {
	testCases := []struct {
		name          string
		queryBuilder  querybuilder.QueryBuilder
		filterParam   []querybuilder.Filter
		seprator      querybuilder.FilterSplitOperator
		expectedRaw   string
		expectedError string
	}{
		{
			name:         "add a filter to query base",
			queryBuilder: querybuilder.New(mainQuery),
			seprator:     querybuilder.And,
			filterParam: []querybuilder.Filter{
				{
					Column: "c.CPF",
					Op:     querybuilder.Equal,
					Value:  "53418727034",
				},
			},
			expectedRaw: fmt.Sprintf("%s \n%s", mainQuery, "WHERE c.CPF = @p1"),
		},
		{
			name:         "add a multi filters to query base",
			queryBuilder: querybuilder.New(mainQuery),
			seprator:     querybuilder.And,
			filterParam: []querybuilder.Filter{
				{
					Column: "c.CPF",
					Op:     querybuilder.Equal,
					Value:  "53418727034",
				},
				{
					Column: "p.EstagioPedido",
					Op:     querybuilder.In,
					Value:  []int{0, 99},
				},
			},
			expectedRaw: fmt.Sprintf("%s \n%s \n%s", mainQuery, "WHERE c.CPF = @p1", "AND p.EstagioPedido IN (@p2,@p3)"),
		},
		{
			name:         "add a filter unexpected value to in operator",
			queryBuilder: querybuilder.New(mainQuery),
			seprator:     querybuilder.And,
			expectedRaw:  fmt.Sprintf("%s \n%s", mainQuery, "WHERE c.CPF = @p1"),
			filterParam: []querybuilder.Filter{
				{
					Column: "c.CPF",
					Op:     querybuilder.Equal,
					Value:  "53418727034",
				},
				{
					Column: "p.EstagioPedido",
					Op:     querybuilder.In,
					Value:  "0,99",
				},
			},
			expectedError: querybuilder.ErrorCreateInOperator.Error(),
		},
		{
			name:         "add empty value to in operator",
			queryBuilder: querybuilder.New(mainQuery),
			seprator:     querybuilder.And,
			expectedRaw:  fmt.Sprintf("%s \n%s", mainQuery, "WHERE c.CPF = @p1"),
			filterParam: []querybuilder.Filter{
				{
					Column: "c.CPF",
					Op:     querybuilder.Equal,
					Value:  "53418727034",
				},
				{
					Column: "p.EstagioPedido",
					Op:     querybuilder.In,
					Value:  []string{},
				},
			},
			expectedError: querybuilder.ErrorCreateInOperator.Error(),
		},
		{
			name:         "add invalid operator",
			queryBuilder: querybuilder.New(mainQuery),
			seprator:     querybuilder.And,
			expectedRaw:  mainQuery,
			filterParam: []querybuilder.Filter{
				{
					Column: "c.CPF",
					Op:     "#",
					Value:  "53418727034",
				},
			},
			expectedError: querybuilder.ErrorInvalidOperator.Error(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// err := tc.queryBuilder.AddFilters(tc.filterParam, tc.seprator)

			raw := tc.queryBuilder.RawQuery()
			if strings.Compare(raw, tc.expectedRaw) != 0 {
				t.Errorf("unexpected raw query given %s\n%s", raw, tc.expectedRaw)
			}
		})
	}
}

func TestAddGroupBy(t *testing.T) {
	testCases := []struct {
		name         string
		queryBuilder querybuilder.QueryBuilder
		filterParam  []querybuilder.Filter
		seprator     querybuilder.FilterSplitOperator
		groupClause  string
		expectedRaw  string
	}{
		{
			name:         "add group by to query base",
			queryBuilder: querybuilder.New(mainQuery),
			seprator:     querybuilder.And,
			expectedRaw:  fmt.Sprintf("%s \n%s \n%s", mainQuery, "WHERE c.CPF = @p1", "GROUP BY c.CPF"),
			filterParam: []querybuilder.Filter{
				{
					Column: "c.CPF",
					Op:     querybuilder.Equal,
					Value:  "53418727034",
				},
			},
			groupClause: "c.CPF",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.queryBuilder.AddFilters(tc.filterParam, tc.seprator)
			if err != nil {
				t.Errorf("unexpected error %s", err.Error())
			}

			tc.queryBuilder.AddGroupBy(tc.groupClause)
			raw := tc.queryBuilder.RawQuery()
			if strings.Compare(raw, tc.expectedRaw) != 0 {
				t.Errorf("unexpected raw query given %s\n%s", raw, tc.expectedRaw)
			}
		})
	}
}
