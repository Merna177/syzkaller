package prog

import (
	"testing"
)

func TestFilterArguments(t *testing.T) {
	tests := []struct {
		os     string
		arch   string
		prog   string
		result bool
	}{
		{
			"linux",
			"amd64",
			`add_key(&(0x7f0000005f40)='dns_resolver\x00', &(0x7f0000005f80)={'syz', 0x0, 0x7a}, &(0x7f0000005f20)="786015083dc3dbe94536578dc260891f45c4b3713a210099", 0x70, 0xffffffffffffffff)`,
			true,
		},
		{
			"linux",
			"amd64",
			`add_key(&(0x7f0000005f40)='dns_resolver\x00', &(0x7f0000005f50)={'syz', 0x0, 0x7a}, &(0x7f0000005f90)="786015083dc3dbe94536578dc260891f45c4b3713a210099", 0x70, 0xffffffffffffffff)`,
			false,
		},
		{
			"test",
			"64",
			"dfetch0(&(0x7f0000000000)='123')",
			false,
		},
		{
			"test",
			"64",
			`dfetch1(&(0x7f0000000040)={0x20000000000002c8, &(0x7f0000000000)=[{'syz'}]})`,
			true,
		},
	}
	t.Parallel()
	for ti, test := range tests {
		target, err := GetTarget(test.os, test.arch)
		if err != nil {
			t.Fatal(err)
		}
		p, err := target.Deserialize([]byte(test.prog), Strict)
		if err != nil {
			t.Fatalf("failed to deserialize the program: %v", err)
		}
		ret := HasOverLappedArgs(p)
		Gooo(p, t)
		if ret == test.result {
			t.Logf("success on test %v", ti)
		} else {
			t.Fatalf("failed on test %v", ti)
		}
	}
}
