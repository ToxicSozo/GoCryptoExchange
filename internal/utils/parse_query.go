package utils

import (
	"fmt"
	"strings"
)

type QueryData struct {
	Command string

	Tables  []string
	Columns []string
	Values  []string

	Condition string
}

func ParseQuery(query string) (*QueryData, error) {
	tokens := strings.Fields(query)
	if len(tokens) == 0 {
		return nil, fmt.Errorf("empty query")
	}

	cmd := &QueryData{Command: strings.ToUpper(tokens[0])}

	switch cmd.Command {
	case "INSERT":
		return parseInsert(cmd, tokens)
	case "SELECT":
		return parseSelect(cmd, tokens)
	case "DELETE":
		return parseDelete(cmd, tokens)
	default:
		return nil, fmt.Errorf("unknown command: %s", cmd.Command)
	}
}

func parseInsert(cmd *QueryData, tokens []string) (*QueryData, error) {
	if len(tokens) < 4 || strings.ToUpper(tokens[1]) != "INTO" || strings.ToUpper(tokens[3]) != "VALUES" {
		return nil, fmt.Errorf("invalid INSERT syntax")
	}

	cmd.Tables = append(cmd.Tables, tokens[2])
	cmd.Values = tokens[4:]
	return cmd, nil
}

func parseSelect(cmd *QueryData, tokens []string) (*QueryData, error) {
	fromIndex := indexOf(tokens, "FROM")
	if fromIndex == -1 || fromIndex < 2 {
		return nil, fmt.Errorf("invalid SELECT syntax")
	}

	cmd.Columns = splitAndTrim(tokens[1:fromIndex], ",")
	whereIndex := indexOf(tokens, "WHERE")

	if whereIndex == -1 {
		cmd.Tables = splitAndTrim(tokens[fromIndex+1:], ",")
	} else {
		cmd.Tables = splitAndTrim(tokens[fromIndex+1:whereIndex], ",")
		cmd.Condition = strings.Join(tokens[whereIndex+1:], " ")
	}
	return cmd, nil
}

func parseDelete(cmd *QueryData, tokens []string) (*QueryData, error) {
	if len(tokens) < 3 || strings.ToUpper(tokens[1]) != "FROM" {
		return nil, fmt.Errorf("invalid DELETE syntax")
	}

	cmd.Tables = append(cmd.Tables, tokens[2])
	whereIndex := indexOf(tokens, "WHERE")

	if whereIndex != -1 {
		cmd.Condition = strings.Join(tokens[whereIndex+1:], " ")
	}
	return cmd, nil
}

func indexOf(slice []string, target string) int {
	for i, v := range slice {
		if strings.ToUpper(v) == target {
			return i
		}
	}
	return -1
}

func splitAndTrim(slice []string, sep string) []string {
	var result []string
	for _, item := range slice {
		trimmed := strings.Trim(item, sep)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
