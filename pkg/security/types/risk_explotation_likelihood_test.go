/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package types

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ParseRiskExploitationLikelihoodTest struct {
	input         string
	expected      RiskExploitationLikelihood
	expectedError error
}

func TestParseRiskExploitationLikelihood(t *testing.T) {
	testCases := map[string]ParseRiskExploitationLikelihoodTest{
		"unlikely": {
			input:    "unlikely",
			expected: Unlikely,
		},
		"likely": {
			input:    "likely",
			expected: Likely,
		},
		"very-likely": {
			input:    "very-likely",
			expected: VeryLikely,
		},
		"frequent": {
			input:    "frequent",
			expected: Frequent,
		},
		"default": {
			input:    "",
			expected: Likely,
		},
		"unknown": {
			input:         "unknown",
			expectedError: errors.New("Unable to parse into type: unknown"),
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			actual, err := ParseRiskExploitationLikelihood(testCase.input)

			assert.Equal(t, testCase.expected, actual)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}
