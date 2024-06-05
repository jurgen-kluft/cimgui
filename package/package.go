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
	mainlib := denv.SetupDefaultCppLibProject("cimgui", "github.com\\jurgen-kluft\\cimgui")
	mainlib.Dependencies = append(mainlib.Dependencies, glfwpkg.GetMainLib())

	if denv.OS == "windows" {
		//mainlib.AddDefine("VK_USE_PLATFORM_WIN32_KHR")
	} else if denv.OS == "darwin" {
		//mainlib.AddDefine("VK_USE_PLATFORM_MACOS_MVK")
	} else if denv.OS == "linux" {
		//mainlib.AddDefine("VK_USE_PLATFORM_XLIB_KHR")
	}

	mainpkg.AddMainLib(mainlib)

	// unittest project
	maintest := denv.SetupDefaultCppTestProject("cimgui"+"test", "github.com\\jurgen-kluft\\cimgui")
	maintest.Dependencies = append(maintest.Dependencies, unittestpkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, glfwpkg.GetMainLib())
	maintest.Dependencies = append(maintest.Dependencies, mainlib)

	mainpkg.AddUnittest(maintest)

	return mainpkg
}
