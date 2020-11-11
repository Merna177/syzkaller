package prog

import (
	"fmt"
	"testing"
)

func TestFilterArguments(t *testing.T) {
	tests := [][2]string{
		{
			`dfetch0(&(0x7f0000000000)='123')`,
			"false",
		},
		{
			`dfetch1(&(0x7f0000000000)={0x0, {0x0}})`,
			"true",
		},
	}
	target := initTargetTest(t, "test", "64")
	for ti, test := range tests {
		t.Run(fmt.Sprint(ti), func(t *testing.T) {
			t.Parallel()
			p, err := target.Deserialize([]byte(test[0]), Strict)
			if err != nil {
				t.Fatalf("failed to deserialize the program: %v", err)
			}
			ret := DFetchAnalysis(p)
			if ret && test[1] == "true" || !ret && test[1] == "false" {
				t.Logf("success on test %v", ti)
			} else {
				t.Fatalf("failed on test %v", ti)
			}
		})
	}
}
