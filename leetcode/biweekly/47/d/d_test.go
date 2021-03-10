// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`4`, `[[1,2],[2,4],[1,3],[2,3],[2,1]]`, `[2,3]`, 
			`[6,5]`,
		},
		{
			`5`, `[[1,5],[1,5],[3,4],[2,5],[1,3],[5,1],[2,3],[2,5]]`, `[1,2,3,4,5]`, 
			`[10,10,9,8,6]`,
		},
		// TODO 测试入参最小的情况
		{
			`8`, `[[3,2],[5,7],[4,2],[6,2],[5,6],[7,5],[5,2],[5,1],[5,4],[7,2],[3,5],[6,2],[6,4],[3,4],[8,6],[4,8],[1,4],[2,7],[7,8]]`, `[7,2,3,3,4,1,4,0,2,1,1,4,0,7,10,0,0]`,
			`[20,28,28,28,28,28,28,28,28,28,28,28,28,20,5,28,28]`,
		},
		{
			`8`, `[[3,2],[5,7],[4,2],[6,2],[5,6],[7,5],[5,2],[5,1],[5,4],[7,2],[3,5],[6,2],[6,4],[3,4],[8,6],[4,8],[1,4],[2,7],[7,8]]`, `[10]`,
			`[5]`,
		},
		{
			`10`, `[[10,7],[7,3],[10,3],[7,4],[9,10],[3,10],[4,5]]`, `[0,3,2,2,3,2,0,1,0,1,1,1,1,0,0,2,1,1,3]`,
			`[39,16,25,25,16,25,39,31,39,31,31,31,31,39,39,25,31,31,16]`,
		},
	}
	targetCaseNum :=  -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, countPairs, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-47/problems/count-pairs-of-nodes/
