package etcd

import (
	"fmt"
	"net/url"
	"reflect"
)

type options map[string]interface{}

// An internally-used data structure that represents a mapping
// between valid options and their kinds
type validOptions map[string]reflect.Kind

// Valid options for GET, PUT, POST, DELETE
// Using CAPITALIZED_UNDERSCORE to emphasize that these
// values are meant to be used as constants.
var (
	VALID_GET_OPTIONS = validOptions{
		"recursive":  reflect.Bool,
		"consistent": reflect.Bool,
		"sorted":     reflect.Bool,
		"wait":       reflect.Bool,
		"waitIndex":  reflect.Uint64,
	}

	VALID_PUT_OPTIONS = validOptions{
		"prevValue": reflect.String,
		"prevIndex": reflect.Uint64,
		"prevExist": reflect.Bool,
	}

	VALID_POST_OPTIONS = validOptions{}

	VALID_DELETE_OPTIONS = validOptions{
		"recursive": reflect.Bool,
	}
)

// Convert options to a string of HTML parameters
func (ops options) toParameters(validOps validOptions) (string, error) {
	p := "?"
	values := url.Values{}

	if ops == nil {
		return "", nil
	}

	for k, v := range ops {
		// Check if the given option is valid (that it exists)
		kind := validOps[k]
		if kind == reflect.Invalid {
			return "", fmt.Errorf("Invalid option: %v", k)
		}

		// Check if the given option is of the valid type
		t := reflect.TypeOf(v)
		if kind != t.Kind() {
			return "", fmt.Errorf("Option %s should be of %v kind, not of %v kind.",
				k, kind, t.Kind())
		}

		values.Set(k, fmt.Sprintf("%v", v))
	}

	p += values.Encode()
	return p, nil
}
