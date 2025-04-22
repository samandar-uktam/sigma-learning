package factory

import "fmt"

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newAk47(), nil
	}

	return nil, fmt.Errorf("Wrong gun type!")
}
