package mockData

// BiologicalSex represents information of an indivual biology.
type BiologicalSex struct {
	Name        string
	Description string
}

var BiologicalSexes map[string]BiologicalSex

func init() {
	report("BiologicalSexes")
	BiologicalSexes = make(map[string]BiologicalSex)
	BiologicalSexes["M"] = BiologicalSex{Name: "Male", Description: ""}
	BiologicalSexes["F"] = BiologicalSex{Name: "Female", Description: ""}
	BiologicalSexes["I"] = BiologicalSex{Name: "Intersex", Description: ""}
	BiologicalSexes["O"] = BiologicalSex{Name: "Other", Description: ""}
}

func GetBiologyList() []string {
	rtn := []string{}
	for k := range BiologicalSexes {
		rtn = append(rtn, k)
	}
	return rtn
}

func GetBiologyInfo(biology string) BiologicalSex {
	return BiologicalSexes[biology]
}

func IsValidBiology(biology string) bool {
	_, ok := BiologicalSexes[biology]
	return ok
}
