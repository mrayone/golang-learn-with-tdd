package tags

import (
	"reflect"
)

type Config struct {
	AllowPlay       string `origin:"settings.canPlay" json:"allowPlay"`
	AllowWatchVideo string `origin:"settings.canWatchVideo" json:"allowWatchVideo"`
	PriceOfMovie    string `origin:"settings.priceMovie" json:"priceOfMovie"`
	ObjectValue     string `origin:"settings.value" json:"objectValue"`
	ToggleDiscount  string `origin:"settings.discountEnabled" json:"toggleDiscount"`
	Message         string `origin:"settings.defaultMessage" json:"message"`
}

func (c Config) GetOrigins() map[string]string {
	st := reflect.TypeOf(c)
	output := make(map[string]string)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		output[field.Tag.Get("origin")] = field.Tag.Get("json")
	}

	return output
}

func FindAllConfig(data map[string]string) (map[string]string, error) {
	configs := make(map[string]string)
	config := Config{}
	for i, v := range config.GetOrigins() {
		if dbv, ok := data[i]; ok {
			configs[v] = dbv
		}
	}

	//jsonStr, err := json.Marshal(configs)
	//if err != nil {
	//	return config, err
	//}
	//
	//if err := json.Unmarshal(jsonStr, &config); err != nil {
	//	return Config{}, err
	//}

	return configs, nil
}

func FindAllWithMaps(data map[string]string, translate map[string]string) map[string]string {
	output := map[string]string{}

	for i, dbv := range data {
		if tl, ok := translate[i]; ok {
			output[tl] = dbv
		}
	}

	return output
}
