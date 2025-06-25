package util

import "strconv"

// GetStringParam returns the specified parameter if provided or the default
// if not.
func GetStringParam(params map[string]string, name, def string) string {
	v, ok := params[name]
	if !ok {
		return def
	}
	return v
}

// GetIntParam returns the specified parameter as an integer if provided or
// the default if not.
func GetIntParam(params map[string]string, name string, def int) int {
	v, ok := params[name]
	if !ok {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return i
}

// GetFloatParam returns the specified parameter as a float if provided or the
// default if not.
func GetFloatParam(params map[string]string, name string, def float64) float64 {
	v, ok := params[name]
	if !ok {
		return def
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		panic(err)
	}
	return f
}
