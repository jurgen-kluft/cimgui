package main

import (
	ccode "github.com/jurgen-kluft/ccode"
	cpkg "github.com/jurgen-kluft/cimgui/package"
)

func main() {
	if ccode.Init() {
		ccode.GenerateClangFormat()
		ccode.GenerateGitIgnore()
		ccode.Generate(cpkg.GetPackage())
	}
}
