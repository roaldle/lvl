/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"bitbucket.org/level27/lvl/types"
	"github.com/Jeffail/gabs/v2"
	"github.com/spf13/cobra"
)

// Contains parameters passed to get commands, like filter and max number of entries.
var optGetParameters types.CommonGetParams

// Add common flags for get-style commands, such as --number and --filter.
func addCommonGetFlags(cmd *cobra.Command) {
	pf := cmd.Flags()

	pf.IntVarP(&optGetParameters.Limit, "number", "n", optGetParameters.Limit, "How many things should we retrieve from the API?")
	pf.StringVarP(&optGetParameters.Filter, "filter", "f", optGetParameters.Filter, "How to filter API results?")
}

//check for valid ID as type INT.
func checkSingleIntID(ids []string, entity string) int {
	id, err := strconv.Atoi(ids[0])
	if err != nil {
		log.Fatalf("Not a valid %v ID!", entity)
	}
	return id
}

// Try to split the given cmd args into ID's (works with whitespace and komma's)
func CheckForMultipleIDs(ids []string) []string {
	var currIds []string

	for _, id := range ids {
		tempId := strings.Split(id, ",")
		currIds = append(currIds, tempId...)
	}

	return currIds
}
// --------------------------- DYNAMICALY SETTING PARAMETERS

// function used for commands with dynamic parameters. (different parameters defined by 1 flag)
func SplitCustomParameters(customP []string) (map[string]interface{}, error) {
	checkedParameters := make(map[string]interface{})
	var err error
	// loop over raw data set by user with -p flag
	for _, setParameter := range customP {
		// check if correct way is used to define parameters -> key=value
		if strings.Contains(setParameter, "=") {

			// split each parameter set by user into its key and value. put them in the dictionary
			line := strings.Split(setParameter, "=")
			// some keys can use multiple values. check if values seperated by comma
			if strings.Contains(line[1], ",") {
				values := strings.Split(line[1], ",")
				//removing spaces from splitted values
				for i, _ := range values {
					values[i] = strings.Trim(values[i], " ")
				}
				// add key value pair to dict
				checkedParameters[strings.Trim(line[0], " ")] = values
			} else {
				// add key value pair to dict
				checkedParameters[strings.Trim(line[0], " ")] = strings.Trim(line[1], " ")
			}

		} else {
			// when there is no '=' in the parameter -> error
			message := fmt.Sprintf("Wrong way of defining parameter is used for: '%v'. (use:[ -p key=value ])", setParameter)
			err = errors.New(message)

		}

	}
	return checkedParameters, err
}

// VALIDATION OF PARAMETER VALUES BASED ON JSON
func isValueValidForParameter(container gabs.Container, value interface{}, currentOS string) (bool, bool) {

	// convert value to string
	valueAsString := fmt.Sprintf("%v", value)
	var option types.CookbookParameterOption
	var isAvailableforSystemOS bool = false

	// unmarshal into struct, to easy and clearly manipulate options data
	err := json.Unmarshal([]byte(container.Search(valueAsString).String()), &option)
	if err != nil {
		log.Fatal(err)
	}

	// loop over all possible OS systems for the chosen value
	// and check if it matches the current system OS
	for _, optionalOS := range option.OperatingSystemVersions {
		if optionalOS.Name == currentOS || currentOS == "" {
		
			isAvailableforSystemOS = true

		}
	}
	if !isAvailableforSystemOS {
		message := fmt.Sprintf("Given Value: '%v' can not be installed on current OS: '%v'.", value, currentOS)
		err = errors.New(message)

		log.Fatal(err)
	}

	return isAvailableforSystemOS, option.Exclusive
}
// Open a file passed as an argument.
// This handles the convention of "-" opening stdin.
func openArgFile(file string) io.ReadCloser {
	if file == "-" {
		return os.Stdin
	} else {
		f, err := os.Open(file)
		cobra.CheckErr(err)
		return f
	}
}

// Load JSON settings from arg-specified file and merge it with override settings from other args.
func loadMergeSettings(fileName string, override map[string]interface{}) map[string]interface{} {
	if fileName == "" {
		return override
	}

	file := openArgFile(fileName)

	defer func(){ cobra.CheckErr(file.Close()) }()

	jsonBytes, err := io.ReadAll(file)
	cobra.CheckErr(err)

	var jsonSettings map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonSettings)
	cobra.CheckErr(err)

	return mergeMaps(jsonSettings, override)
}

func mergeMaps(base map[string]interface{}, override map[string]interface{}) map[string]interface{} {
	var newMap = map[string]interface{}{}

	for k, v := range base {
		newMap[k] = v
	}

	for k, v := range override {
		newMap[k] = v
	}

	return newMap
}

// Add a string setting flag to a command, that will be stored in a map.
// This is intended to be easily used with PATCH APIs.
func settingString(c *cobra.Command, settings map[string]interface{}, name string, usage string) {
	c.Flags().Var(&stringMapValue{Map: settings, Name: name}, name, usage)
}

// Add a string setting flag to a command, that will be stored in a map. Shorthand version.
// This is intended to be easily used with PATCH APIs.
func settingStringP(c *cobra.Command, settings map[string]interface{}, name string, short string, usage string) {
	c.Flags().VarP(&stringMapValue{Map: settings, Name: name}, name, short, usage)
}

// Add an int setting flag to a command, that will be stored in a map.
// This is intended to be easily used with PATCH APIs.
func settingInt(c *cobra.Command, settings map[string]interface{}, name string, usage string) {
	c.Flags().Var(&intMapValue{Map: settings, Name: name}, name, usage)
}

// Add an int setting flag to a command, that will be stored in a map. Shorthand version.
// This is intended to be easily used with PATCH APIs.
func settingIntP(c *cobra.Command, settings map[string]interface{}, name string, short string, usage string) {
	c.Flags().VarP(&intMapValue{Map: settings, Name: name}, name, short, usage)
}

// Add a bool setting flag to a command, that will be stored in a map.
// This is intended to be easily used with PATCH APIs.
func settingBool(c *cobra.Command, settings map[string]interface{}, name string, usage string) {
	c.Flags().Var(&boolMapValue{Map: settings, Name: name}, name, usage)
}

// Add a bool setting flag to a command, that will be stored in a map. Shorthand version.
// This is intended to be easily used with PATCH APIs.
func settingBoolP(c *cobra.Command, settings map[string]interface{}, name string, short string, usage string) {
	c.Flags().VarP(&boolMapValue{Map: settings, Name: name}, name, short, usage)
}

// Types below implement the pflag.Value interface for string/int/bool so that they can receive values assigned via command line.

type stringMapValue struct {
	Map  map[string]interface{}
	Name string
}

func (c *stringMapValue) String() string {
	val := c.Map[c.Name]
	if val == nil {
		return ""
	}

	return val.(string)
}

func (c *stringMapValue) Set(val string) error {
	c.Map[c.Name] = val
	return nil
}

func (c *stringMapValue) Type() string {
	return "string"
}

type intMapValue struct {
	Map  map[string]interface{}
	Name string
}

func (c *intMapValue) String() string {
	val := c.Map[c.Name]
	if val == nil {
		return ""
	}

	return strconv.Itoa(val.(int))
}

func (c *intMapValue) Set(val string) error {
	i, err := strconv.Atoi(val)
	if err != nil {
		return err
	}

	c.Map[c.Name] = i
	return nil
}

func (c *intMapValue) Type() string {
	return "int"
}

type boolMapValue struct {
	Map  map[string]interface{}
	Name string
}

func (c *boolMapValue) String() string {
	val := c.Map[c.Name]
	if val == nil {
		return ""
	}

	return strconv.FormatBool(val.(bool))
}

func (c *boolMapValue) Set(val string) error {
	b, err := strconv.ParseBool(val)
	if err != nil {
		return err
	}

	c.Map[c.Name] = b
	return nil
}

func (c *boolMapValue) Type() string {
	return "bool"
}
