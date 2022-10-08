package core

/*	License: GPLv3
	Authors:
		Mirko Brombin <send@mirko.pm>
		Pietro di Caprio <pietro@fabricators.ltd>
	Copyright: 2022
	Description: Apx is a wrapper around apt to make it works inside a container
	from outside, directly on the host.
*/

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	err := CheckContainerTools()
	if err != nil {
		fmt.Println(`One or more core components are not available. 
Please refer to our documentation at https://documentation.vanillaos.org/`)
		panic(err)
	}
}

func CheckContainerTools() error {
	distrobox := exec.Command("which", "distrobox")
	docker := exec.Command("which", "docker")
	podman := exec.Command("which", "podman")

	if err := distrobox.Run(); err != nil {
		err := InstallDistrobox()
		if err != nil {
			return err
		}
	}

	if err := docker.Run(); err != nil {
		if err := podman.Run(); err != nil {
			InstallPodman()
		}
	}

	return nil
}

func InstallDistrobox() error {
	fmt.Println(`Distrobox is not installed. Would you like to install it now? [y/N]`)
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		return err
	}

	if input != "y" {
		fmt.Println("Please install Distrobox in order to use apx!")
		os.Exit(1)
	}

	installDistroboxScript := exec.Command("sudo", "curl", "-s", "https://raw.githubusercontent.com/89luca89/distrobox/main/install | sudo sh")
	err = installDistroboxScript.Run()
	if err != nil {
		fmt.Println("Cannot automatically install distrobox. Please install it manually.")
		return err
	}

	return nil

}

func InstallPodman() error {
	fmt.Println(`Podman is not installed. Would you like to install it now? [y/N]`)
	var input string
	_, err := fmt.Scanln(&input)

	if err != nil {
		return err
	}

	if input != "y" {
		fmt.Println("Please install Podman in order to user apx!")
		os.Exit(1)
	}

	installDistroboxScript := exec.Command("sudo", "apt", "install", "-y", "podman")
	installDistroboxScript.Run()
	if err != nil {
		fmt.Println("Cannot automatically install podman. Please install it manually.")
		return err
	}

	return nil
}
