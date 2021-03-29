package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Conf struct {
		File              string `yaml:"file"`
		NumberAnswer      int    `yaml:"numberAnswer"`
		AnswerPerQuestion int    `yaml:"answerPerQuestion"`
		QuestionsPerGame  int    `yaml:"questionsPerGame"`
		QuestionsUrl      string `yaml:"questionsUrl"`
		AnswerUrl         string `yaml:"answerUrl"`
	}
}

var C *Conf
var err error

func init() {
	C, err = readConf("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

}

// Read the yaml  Config
func readConf(filename string) (*Conf, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &Conf{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}
	return c, nil
}
