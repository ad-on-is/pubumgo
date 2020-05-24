package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	. "github.com/logrusorgru/aurora"
)

func main() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	input, err := ioutil.ReadFile(path + "/pubspec.yaml")
	if err != nil {
		fmt.Println(Red("Could not find pubspec.yaml"))
		return
	}

	bumps := [5]string{"major", "minor", "patch", "build"}
	bump := ""
	keepBuild := false

	if len(os.Args) < 2 {
		fmt.Println("Plese provide one of the bumps: ", Yellow(bumps))
		return
	} else {
		bump = os.Args[1]
		if !contains(bumps, bump) {
			fmt.Println("Unknown bump! Only these are valid: ", Yellow(bumps))
			return
		}
		if len(os.Args) > 2 && os.Args[2] == "-b" {
			keepBuild = true
		}
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.HasPrefix(line, "version: ") {
			lines[i] = bumpVersion(lines[i], bump, keepBuild)
			break
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("pubspec.yaml", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func contains(arr [5]string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func bumpVersion(v string, bump string, keepBuild bool) string {

	regex := regexp.MustCompile(`[\r\n]+`) // windows newline fix
	version := strings.Join(strings.Split(regex.ReplaceAllString(v, ""), "version: "), "")

	checkBuild := strings.Split(version, "+")
	hasBuild := len(checkBuild) == 2

	numbers := strings.Split(checkBuild[0], ".")
	vMajor, err := strconv.Atoi(numbers[0])
	vMinor, err := strconv.Atoi(numbers[1])
	vPatch, err := strconv.Atoi(numbers[2])
	vBuild := 0
	if hasBuild {
		vBuild, err = strconv.Atoi(checkBuild[1])
	}

	if err != nil {
		fmt.Println(err)
	}

	switch bump {
	case "major":
		vMajor++
		vMinor = 0
		vPatch = 0
		if !keepBuild {
			vBuild = 0
			hasBuild = false
		}

		break
	case "minor":
		vMinor++
		vPatch = 0
		if !keepBuild {
			vBuild = 0
			hasBuild = false
		}
		break
	case "patch":
		vPatch++
		if !keepBuild {
			vBuild = 0
			hasBuild = false
		}
		break
	case "build":
		vBuild++
		break
	}

	maj := strconv.Itoa(vMajor)
	min := strconv.Itoa(vMinor)
	patch := strconv.Itoa(vPatch)
	build := "+" + strconv.Itoa(vBuild)

	if !hasBuild && bump != "build" {
		build = ""
	}

	newVersion := maj + "." + min + "." + patch + build
	fmt.Println("pubspec.yml bumped successfully from", Yellow(version), "->", Green(newVersion))
	return "version: " + newVersion
}
