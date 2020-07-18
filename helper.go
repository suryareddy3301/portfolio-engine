package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

func CreateConfig() (*Server, error) {
	server := &Server{
		About:          "Software Developer with n years of experience ...",
		BasicInfo:      newBasicInfo(),
		Skills:         getSkills(),
		SocialNetworks: getSocilaNetworks(),
		Education:      getEducation(),
		Experience:     newExperience(),
		Theme:          getTheme(),
	}

	file, err := os.OpenFile("config.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	b, err := yaml.Marshal(server)
	if err != nil {
		return nil, err
	}
	_, err = file.Write(b)
	if err != nil {
		return nil, err
	}
	return server, nil
}

func getTheme() *Theme {
	return &Theme{
		Name:            "main",
		BackgroundImage: "cc-bg-11.jpg",
		ProfileImage:    "SURYA.jpg",
	}
}

func getSkills() []*kv {
	entries := []*kv{
		&kv{Key: "golang", Value: "70%"},
		&kv{Key: "javascript", Value: "50%"},
		&kv{Key: "HTML", Value: "60%"},
		&kv{Key: "jenkins", Value: "50%"},
		&kv{Key: "drone.io", Value: "50%"},
		&kv{Key: "git", Value: "75%"},
		&kv{Key: "Agile and Scrum", Value: "75%"},
	}
	return entries
}

func getSocilaNetworks() []*kv {
	return []*kv{
		&kv{
			Key:   "github",
			Value: "https://github.com",
		},
		&kv{
			Key:   "linkedin",
			Value: "https://linkedin.com",
		},
		&kv{
			Key:   "instagram",
			Value: "https://instagram.com",
		}, &kv{
			Key:   "medium",
			Value: "https://medium.com",
		},
	}
}

func newBasicInfo() *BasicInfo {
	details := []*kv{
		&kv{Key: "Email", Value: "youremail@domain.com"},
		&kv{Key: "Phone", Value: "(+00) 123456789"},
		&kv{Key: "Address", Value: "City, State ZIP"},
	}
	basicInfo := &BasicInfo{
		Name:    "Your Name",
		Details: details,
	}
	return basicInfo
}

func newExperience() []*Experience {
	return []*Experience{
		&Experience{
			Company:     "Company 1",
			Date:        "2020 - Present",
			Description: "Euismod massa scelerisque suspendisse fermentum habitant vitae ullamcorper magna quam iaculis, tristique sapien taciti mollis interdum sagittis libero nunc inceptos tellus, hendrerit vel eleifend primis lectus quisque cubilia sed mauris. Lacinia porta vestibulum diam integer quisque eros pulvinar curae, curabitur feugiat arcu vivamus parturient aliquet laoreet at, eu etiam pretium molestie ultricies sollicitudin dui.",
			Designation: "Project Lead",
		},
		&Experience{
			Company:     "Company 2",
			Date:        "2016 - 2020",
			Description: "Euismod massa scelerisque suspendisse fermentum habitant vitae ullamcorper magna quam iaculis, tristique sapien taciti mollis interdum sagittis libero nunc inceptos tellus, hendrerit vel eleifend primis lectus quisque cubilia sed mauris. Lacinia porta vestibulum diam integer quisque eros pulvinar curae, curabitur feugiat arcu vivamus parturient aliquet laoreet at, eu etiam pretium molestie ultricies sollicitudin dui.",
			Designation: "Senior Principal Development Specialist",
		},
		&Experience{
			Company:     "Company 3",
			Date:        "2013 - 2016",
			Description: "Euismod massa scelerisque suspendisse fermentum habitant vitae ullamcorper magna quam iaculis, tristique sapien taciti mollis interdum sagittis libero nunc inceptos tellus, hendrerit vel eleifend primis lectus quisque cubilia sed mauris. Lacinia porta vestibulum diam integer quisque eros pulvinar curae, curabitur feugiat arcu vivamus parturient aliquet laoreet at, eu etiam pretium molestie ultricies sollicitudin dui.",
			Designation: "Senior Software Engineer",
		},
	}
}

func getEducation() []*Education {
	return []*Education{
		&Education{
			Date:            "2008 - 2012",
			Degree:          "Bachelor's Degree",
			InstitutionName: "Your University Name",
			Speciality:      "Bachelor of Information Technology",
			Description:     "Habitasse venenatis commodo tempor eleifend arcu sociis sollicitudin ante pulvinar ad, est porta cras erat ullamcorper volutpat metus duis platea convallis, tortor primis ac quisque etiam luctus nisl nullam fames. Ligula purus suscipit tempus nascetur curabitur donec nam ullamcorper, laoreet nullam mauris dui aptent facilisis neque elementum ac, risus semper felis parturient fringilla rhoncus eleifend.",
		},
	}
}
