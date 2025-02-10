package dao

import "github.com/mt1976/frantic-core/logger"

var name = "Lookup"

type Lookup struct {
	Data []LookupData
}

type LookupData struct {
	Key          string
	Value        string
	AltID        string
	Description  string
	ObjectDomain string
	Selected     bool
}

func (a *Lookup) Spew() error {
	// Spew the Audit Data
	noAudit := len(a.Data)
	//logger.InfoLogger.Printf(" No Updates=[%v]", noAudit)
	if noAudit > 0 {
		for i := 0; i < noAudit; i++ {
			upd := a.Data[i]
			logger.TraceLogger.Printf("LookupData[%v] Key=[%v] Value=[%v]", i, upd.Key, upd.Value)
		}
	}
	return nil
}
