package routeSet

import (
	"regexp"
	"strconv"
)

type ParameterType string

const (
	NumberParameter = ParameterType("num")
	StringParameter = ParameterType("string")
)

var (
	urlRegexp = regexp.MustCompile(`^[A-Za-z0-9\-\._~]+$`)
)

type routePart struct {
	// if strict is True, then this part of the route should be strictly equal to strictPart field
	// if strict is False, then this part of the route is a parameter
	Strict     bool
	strictPart string

	parameterType ParameterType
}

// NewRoutePart creates a part of the route
// part shoud be or a string of unreserved URL symbols, so it will be strict
// or part can be a parameter: {num}, {string}
func NewRoutePart(part string) (routePart, error) {
	if part[0] == '{' && part[len(part)-1] == '}' {
		center := part[1 : len(part)-1]

		switch center {
		case "num":
			return routePart{
				Strict:        false,
				parameterType: NumberParameter,
			}, nil
		case "string":
			return routePart{
				Strict:        false,
				parameterType: StringParameter,
			}, nil
		default:
			return routePart{}, NewErrorBadRoutePart("part wrapped with {} perceived parameter part, and shoud be or {num} or {string}")
		}
	} else if urlRegexp.MatchString(part) {
		return routePart{
			Strict:     true,
			strictPart: part,
		}, nil
	} else {
		return routePart{}, NewErrorBadRoutePart("part not wrapped with {} perceived strict, and should contain only URL unreserved symbols")
	}
}

// Compare compares a given part of url with actual routePart
// returns false if given part doesn't match and true if matches
// second return is parameter if it was in routePart, check for nil
// if second return value is not nil, so it is parameter
func (p *routePart) Compare(part string) (bool, any) {

	if p.Strict {
		return part == p.strictPart, nil

	} else if p.parameterType == NumberParameter {
		num, err := strconv.Atoi(part)
		if err != nil {
			return false, nil
		}

		return true, num
	} else if p.parameterType == StringParameter {
		return true, part

	} else {
		panic("unreal route part")
	}
}
