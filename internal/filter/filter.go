package filter

import (
	"fmt"
	"regexp"
	"strings"
)

type Operator int

const (
	NOTDEFINED Operator = iota
	EQ
	NE
	GT
	LT
	GE
	LE
)

type Filter struct {
	Value    string
	Operator Operator
}

func NewFilter(operatorStr, value string) (Filter, error) {
	var operador Operator
	switch operatorStr {
	case "==":
		operador = EQ
	case "=":
		operador = EQ
	case "!=":
		operador = NE
	case ">":
		operador = GT
	case "<":
		operador = LT
	case ">=":
		operador = GE
	case "<=":
		operador = LE
	default:
		return Filter{}, fmt.Errorf("Operator '%s' not supported", operador)
	}
	return Filter{Operator: operador, Value: value}, nil
}

func GetFilters(header, filterString string) (map[string]Filter, error) {
	filters := make(map[string]Filter)

	if len(filterString) == 0 {
		return nil, nil
	}

	re := regexp.MustCompile(`(\w+)\s*(==|!=|>=|<=|>|<|=)\s*(\w+)`)

	for _, line := range strings.Split(filterString, "\n") {
		matches := re.FindStringSubmatch(line)
		if len(matches) != 4 {
			return nil, fmt.Errorf("Invalid filter: '%s'", line)
		}
		colName := matches[1]
		if !strings.Contains(header, matches[1]) {
			return nil, fmt.Errorf("Header '%s' not found in CSV file/string", colName)
		}

		filter, err := NewFilter(matches[2], matches[3])
		if err != nil {
			return nil, err
		}

		filters[colName] = filter
	}
	return filters, nil
}

func getOperatorConst(op string) (Operator, error) {
	switch op {
	case "=":
		return EQ, nil
	case "!=":
		return NE, nil
	case ">":
		return GT, nil
	case "<":
		return LT, nil
	case ">=":
		return GE, nil
	case "<=":
		return LE, nil
	default:
		return NOTDEFINED, fmt.Errorf("Filter: error operator not defined")
	}
}

func (F *Filter) Filtrar(comparando string) bool {
	result := strings.Compare(comparando, F.Value)
	switch F.Operator {
	case EQ:
		return result == 0
	case GT:
		return result > 0
	case LT:
		return result < 0
	case NE:
		return result != 0
	case GE:
		return result >= 0
	case LE:
		return result <= 0
	default:
		return true
	}
}
