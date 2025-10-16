package cimgui

import (
	"github.com/jurgen-kluft/ccode/denv"
	cglfw "github.com/jurgen-kluft/cglfw/package"
	cunittest "github.com/jurgen-kluft/cunittest/package"
)

const (
	repo_path = "github.com\\jurgen-kluft"
	repo_name = "cimgui"
)

func GetPackage() *denv.Package {
	name := repo_name

	// dependencies
	cunittestpkg := cunittest.GetPackage()
	glfwpkg := cglfw.GetPackage()

	// main package
	mainpkg := denv.NewPackage(repo_path, repo_name)
	mainpkg.AddPackage(cunittestpkg)
	mainpkg.AddPackage(glfwpkg)

	// main library
	mainlib := denv.SetupCppLibProject(mainpkg, name)
	mainlib.AddDependencies(glfwpkg.GetMainLib())

	if denv.IsWindows() {
		//mainlib.AddDefine("VK_USE_PLATFORM_WIN32_KHR")
	} else if denv.IsMacOS() {
		//mainlib.AddDefine("VK_USE_PLATFORM_MACOS_MVK")
	} else if denv.IsLinux() {
		//mainlib.AddDefine("VK_USE_PLATFORM_XLIB_KHR")
	}

	// test library
	testlib := denv.SetupCppTestLibProject(mainpkg, name)
	testlib.AddDependencies(glfwpkg.GetTestLib())
	testlib.AddDependencies(cunittestpkg.GetTestLib())

	// unittest project
	maintest := denv.SetupCppTestProject(mainpkg, name)
	maintest.AddDependencies(cunittestpkg.GetMainLib())
	maintest.AddDependency(testlib)

	mainpkg.AddMainLib(mainlib)
	mainpkg.AddTestLib(testlib)
	mainpkg.AddUnittest(maintest)
	return mainpkg
}
