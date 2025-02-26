package mockData

// Title represents information of an indivual title.
type Title struct {
	Title string // The title
}

var Titles map[string]Title

func init() {
	report("Titles")
	Titles = make(map[string]Title)
	Titles["Mr"] = Title{Title: "Mr"}
	Titles["Mrs"] = Title{Title: "Mrs"}
	Titles["Miss"] = Title{Title: "Miss"}
	Titles["Ms"] = Title{Title: "Ms"}
	Titles["Dr"] = Title{Title: "Dr"}
	Titles["Prof"] = Title{Title: "Prof"}
}

func GetList() []string {
	rtn := []string{}
	for k := range Titles {
		rtn = append(rtn, k)
	}
	return rtn
}

func IsValidTitle(in string) bool {
	_, ok := Titles[in]
	return ok
}
