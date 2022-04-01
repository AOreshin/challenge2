package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input, expected []int
}

func TestToTrunks(t *testing.T) {
	tests := []testCase{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 0}, []int{2}},
		{[]int{1, 0, 0}, []int{3}},
		{[]int{1, 0, 1, 1}, []int{2, 1, 1}},
	}
	for _, test := range tests {
		t.Run(name(test), func(t *testing.T) {
			result := toTrunks(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestCut(t *testing.T) {
	tests := []testCase{
		{[]int{}, []int{}},
		{[]int{1, 1}, []int{1, 1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{2, 1}},
		{[]int{3}, []int{3}},
		{[]int{2, 2}, []int{2, 1, 1}},
		{[]int{4}, []int{3, 1}},
		{[]int{5}, []int{3, 2}},
		{[]int{6}, []int{3, 3}},
	}
	for _, test := range tests {
		t.Run(name(test), func(t *testing.T) {
			result := cut(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestCalcProfit(t *testing.T) {
	type testCase struct {
		trunks []int
		price  int
	}
	tests := []testCase{
		{[]int{1}, -1},
		{[]int{2}, 3},
		{[]int{3}, 1},
		{[]int{1, 2, 3}, 3},
		{[]int{2, 2, 3}, 7},
	}
	for _, test := range tests {
		name := fmt.Sprintf("trunks=%v price=%d", test.trunks, test.price)
		t.Run(name, func(t *testing.T) {
			price := calcProfit(test.trunks)
			assert.Equal(t, test.price, price)
		})
	}
}

func TestBestCut(t *testing.T) {
	type testCase struct {
		permutationSource []int
		bestCuts          [][]int
		price             int
	}
	tests := []testCase{
		{
			[]int{1, 2, 3},
			[][]int{
				{1, 3, 2},
				{2, 3, 1},
			},
			4,
		},
		{
			[]int{1, 2, 1},
			[][]int{
				{1, 2, 1},
				{2, 1, 1},
			},
			1,
		},
		{
			[]int{1, 2},
			[][]int{
				{1, 2},
				{2, 1},
			},
			2,
		},
		{
			[]int{1, 4},
			[][]int{
				{1, 4},
			},
			5,
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("source=%v best cuts=%v price=%d", test.permutationSource, test.bestCuts, test.price)
		t.Run(name, func(t *testing.T) {
			perms := allPerm(test.permutationSource)
			bestCuts, price := bestCut(perms)
			assert.Equal(t, test.price, price)
			assert.Equal(t, test.bestCuts, bestCuts)
		})
	}
}

func TestCalculate(t *testing.T) {
	type testCase struct {
		cases  [][]int
		orders [][][]int
		price  int
	}
	tests := []testCase{
		{
			[][]int{
				{2, 3, 1},
			},
			[][][]int{
				{
					{2, 3, 1},
					{1, 3, 2},
				},
			},
			4,
		},
		{
			[][]int{
				{1, 2, 1},
				{1, 2},
				{1, 4},
			},
			[][][]int{
				{
					{1, 2, 1},
					{2, 1, 1},
				},
				{
					{1, 2},
					{2, 1},
				},
				{
					{1, 4},
				},
			},
			8,
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("cases=%v price=%d", test.cases, test.price)
		t.Run(name, func(t *testing.T) {
			price, orders := calculate(test.cases)
			assert.Equal(t, test.price, price)
			assert.ElementsMatch(t, test.orders, orders)
		})
	}
}

func TestReadDataEmpty(t *testing.T) {
	buffer := bytes.Buffer{}
	buffer.Write([]byte("0"))
	data, err := readData(&buffer)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(data))
}

func TestReadDataNegativeCases(t *testing.T) {
	tests := []string{
		"-1",
		"a",
		"3\n",
		"1\na",
		"1\n1 a\n",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			buffer := bytes.Buffer{}
			buffer.Write([]byte(test))
			data, err := readData(&buffer)
			assert.Nil(t, data)
			assert.Error(t, err)
		})
	}
}

func TestReadDataPanics(t *testing.T) {
	buffer := bytes.Buffer{}
	buffer.Write([]byte("1\n2\n"))
	assert.Panics(t, func() { readData(&buffer) })
}

func TestReadDataSuccess(t *testing.T) {
	type testCase struct {
		input    string
		expected [][][]int
	}
	tests := []testCase{
		{
			"1\n3 2 3 1\n3\n3 1 2 1\n2 1 2\n2 1 4\n0",
			[][][]int{
				{{2, 3, 1}},
				{{1, 2, 1}, {1, 2}, {1, 4}},
			},
		},
		{
			"1\n3 2 3 1 3 2 1\n0",
			[][][]int{
				{{2, 3, 1}},
			},
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("input=%s expected=%v\n", test.input, test.expected)
		t.Run(name, func(t *testing.T) {
			buffer := bytes.Buffer{}
			buffer.Write([]byte(test.input))
			data, err := readData(&buffer)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, data)
		})
	}
}

func name(test testCase) string {
	return fmt.Sprintf("input=%v expected=%v", test.input, test.expected)
}
