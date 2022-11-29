package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	day int
	part int
	data input
	extraParams []any
}

func newConfig(day int, part int, data input, extraParams ...any) *Config {
	return &Config{day, part, data, extraParams}
}

func ParseConfig() *Config {
	var day, part int
	var inputData string
	flag.IntVar(&day, "d", 1, "the day of the challenge")
	flag.IntVar(&day, "day", 1, "the day of the challenge")
	flag.IntVar(&part, "p", 1, "the part of the challenge")
	flag.IntVar(&part, "part", 1, "the part of the challenge")
	flag.StringVar(&inputData, "i", "real", "the input data to use")
	flag.StringVar(&inputData, "input", "real", "the input data to use")
	flag.Parse()
	fmt.Println(os.Args)
	var data input
	if day < 1 || day > 25 {
		fmt.Println("Day must be between 1 and 25")
	}
	if part < 1 || part > 2 {
		fmt.Println("Part must be 1 or 2")
		return nil
	}
	if inputData == "real" {
		data = *newRealData()
	} else {
		testNumber, err := strconv.Atoi(inputData)
		if err != nil {
			fmt.Println("Input must be either 'real' or a number")
			return nil
		} else {
			data = *newTestData(testNumber)
		}
	}
	args := flag.Args()
	extraParams := make([]any, len(args))
	for i, arg := range args {
		extraParams[i] = arg
	}
	return newConfig(day, part, data, extraParams...)
}

func (c *Config) GetDay() int {
	return c.day
}

func (c *Config) GetExtraParams() []any {
	return c.extraParams
}

func (c *Config) GetInput() string {
	if c.data.real {
		return "real"
	} else {
		return "test" + strconv.Itoa(c.data.test)
	}
}

func (c *Config) GetPart() int {
	return c.part
}

func (c *Config) GetInputFilename() string {
	return fmt.Sprintf("inputs/day%d/data%s.txt", c.day, c.data.String())
}