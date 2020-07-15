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
		&kv{Key: "Email", Value: "devcharmander@gmail.com"},
		&kv{Key: "Phone", Value: "(+91) 7760877384"},
		&kv{Key: "Language", Value: "English, Hindi, Telugu, Odiya"},
		&kv{Key: "Address", Value: "Bangalore, Karnataka 560102"},
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
		&kv{Key: "golang", Value: 70},
		&kv{Key: "javascript", Value: 50},
		&kv{Key: "HTML", Value: 60},
		&kv{Key: "jenkins", Value: 50},
		&kv{Key: "drone.io", Value: 50},
		&kv{Key: "git", Value: 75},
		&kv{Key: "Agile and Scrum", Value: 75},
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
