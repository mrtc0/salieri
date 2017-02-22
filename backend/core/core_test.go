package core

import (
	"testing"
)

func TestContainerPushFile(t *testing.T) {
	err := CodePush("jessie2", "hello", "clang")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestCompile(t *testing.T) {
	err := CodePush("jessie2", "#include <stdio.h>\nint main() {\n    printf(\"HELLO\\n\");\n    return 0;\n}\n", "clang")
	if err != nil {
		t.Error("Transfer Error on TestCompile")
	}
	result := Compile("jessie2", "clang", "")
	if result["stdout"] != "HELLO\n" {
		t.Error("Stdout is not excepted")
	}
}
