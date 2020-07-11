package main

import "github.com/julienschmidt/httprouter"

type Server struct {
	Router         *httprouter.Router
	About          string
	Skills         []*kv
	Experience     []*Experience
	BasicInfo      *BasicInfo
	SocialNetworks map[string]string
	Education      map[string]string
	//	Assets         *Assets
	//	AssetsUrl      string
}

type kv struct {
	key   interface{}
	value interface{}
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

// type Assets struct {
// 	Photo      string
// 	Background string
// 	Resume     string
// }

// func NewAssets() *Assets {
// 	return &Assets{
// 		Background: "/etc/portfolio/assets/images/cc-bg-1.jpg",
// 		Photo:      "/etc/portfolio/assets/images/surya.jpg",
// 		Resume:     "/etc/portfolio/assets/resume.docx",
// 	}
// }

func NewExperience() *Experience {
	return &Experience{
		Company:     "DevOn(Netherlands)",
		Date:        "2016-2020",
		Description: "",
		Designation: "Senior Principal Development Specialist",
	}
}
func NewBasicInfo() *BasicInfo {
	details := []*kv{
		&kv{key: "Address", value: "Bangalore, Karnataka 560102"},
		&kv{key: "Email", value: "devcharmander@gmail.com"},
		&kv{key: "Language", value: "English Hindi Telugu Odiya"},
		&kv{key: "Phone", value: "(+91) 7760877384"},
		&kv{key: "Name", value: "Surya Reddy"},
	}
	basicInfo := &BasicInfo{
		Name:    "Surya Reddy",
		Details: details,
	}
	return basicInfo
}

func NewServer() *Server {
	server := &Server{
		About: "Software Developer with 7 years of experience blah blah",
		//Assets:         NewAssets(),
		BasicInfo:      NewBasicInfo(),
		Skills:         GetSkills(),
		SocialNetworks: GetSocilaNetworks(),
		Education:      make(map[string]string),
		Experience:     []*Experience{NewExperience()},
		//AssetsUrl:      "/etc/portfolio/assets/",
	}
	return server
}

//TODO all these helper methods should read from a file and construct
func GetSkills() []*kv {
	entries := []*kv{
		&kv{key: "golang", value: 70},
		&kv{key: "javascript", value: 50},
		&kv{key: "HTML", value: 60},
		&kv{key: "jenkins", value: 50},
		&kv{key: "drone.io", value: 50},
		&kv{key: "git", value: 75},
		&kv{key: "Agile and Scrum", value: 75},
	}
	return entries
}

func GetSocilaNetworks() map[string]string {
	sns := make(map[string]string)
	sns["Github"] = "https://github.com/suryareddy3301"
	sns["Linkedin"] = "https://linkedin.com"
	sns["Instagram"] = "https://instagram.com"
	sns["Medium"] = "https://medium.com"
	return sns
}
func GetEducation() map[string]string {
	education := make(map[string]string)
	education["2008-2012"] = "Bachelors of Technology in Computer Science"
	return education
}
