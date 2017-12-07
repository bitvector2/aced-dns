package named

import "fmt"

type ResourceRecord struct {
	name        string
	ttl         int
	recordClass string
	recordType  string
	recordData  string
}

func NewResourceRecord(name string, ttl int, recordClass string, recordType string, recordData string) *ResourceRecord {
	if !(recordClass == "IN" || recordClass == "CHAOS") {
		return nil
	}

	return &ResourceRecord{
		name:        name,
		ttl:         ttl,
		recordClass: recordClass,
		recordType:  recordType,
		recordData:  recordData,
	}
}

func (r *ResourceRecord) String() string {
	if r.ttl <= 0 {
		return fmt.Sprintf("%s %s %s %s\n", r.name, r.recordClass, r.recordType, r.recordData)
	}
	return fmt.Sprintf("%s %d %s %s %s\n", r.name, r.ttl, r.recordClass, r.recordType, r.recordData)
}
