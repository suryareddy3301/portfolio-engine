package main

import (
	"io/ioutil"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/yaml.v2"
)

//TODO Logger implementation

type Server struct {
	Router         *httprouter.Router `yaml:"router,omitempty"`
	ProfileImage   string
	About          string
	Skills         []*kv
	Experience     []*Experience
	BasicInfo      *BasicInfo
	SocialNetworks []*kv
	Education      []*Education
}

type kv struct {
	Key   interface{}
	Value interface{}
}

type BasicInfo struct {
	Name    string
	Details []*kv
}

type Experience struct {
	Date        string
	Company     string
	Designation string
	Description string
}

type Education struct {
	Date            string
	InstitutionName string
	Speciality      string // Bachelors in Computer Science Engineering etc
	Degree          string //e.g. Masters bachelors High school etc
	Description     string
}

func NewServer() (*Server, error) {
	server := &Server{}
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal(data, server)
	return server, err
}
