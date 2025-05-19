package cimgui

import (
	"github.com/jurgen-kluft/ccode/denv"
	cglfw "github.com/jurgen-kluft/cglfw/package"
	cunittest "github.com/jurgen-kluft/cunittest/package"
)

// GetPackage returns the package object of 'cimgui'
func GetPackage() *denv.Package {
	// Dependencies
	unittestpkg := cunittest.GetPackage()
	glfwpkg := cglfw.GetPackage()

	// The main (cimgui) package
	mainpkg := denv.NewPackage("cimgui")
	mainpkg.AddPackage(unittestpkg)
	mainpkg.AddPackage(glfwpkg)

	// 'cimgui' library
	mainlib := denv.SetupCppLibProject("cimgui", "github.com\\jurgen-kluft\\cimgui")
	mainlib.AddDependencies(glfwpkg.GetMainLib()...)

	if denv.IsWindows() {
		//mainlib.AddDefine("VK_USE_PLATFORM_WIN32_KHR")
	} else if denv.IsMacOS() {
		//mainlib.AddDefine("VK_USE_PLATFORM_MACOS_MVK")
	} else if denv.IsLinux() {
		//mainlib.AddDefine("VK_USE_PLATFORM_XLIB_KHR")
	}

	mainpkg.AddMainLib(mainlib)

	// unittest project
	maintest := denv.SetupDefaultCppTestProject("cimgui"+"test", "github.com\\jurgen-kluft\\cimgui")
	maintest.AddDependencies(unittestpkg.GetMainLib()...)
	maintest.AddDependencies(glfwpkg.GetMainLib()...)
	maintest.Dependencies = append(maintest.Dependencies, mainlib)

	mainpkg.AddUnittest(maintest)

	return mainpkg
}
