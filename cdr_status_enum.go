// Code generated by go-enum
// DO NOT EDIT!

package qlcchain

import (
	"fmt"
	"strings"
)

const (
	// SettlementStatusUnknown is a SettlementStatus of type Unknown
	SettlementStatusUnknown SettlementStatus = iota
	// SettlementStatusStage1 is a SettlementStatus of type Stage1
	SettlementStatusStage1
	// SettlementStatusSuccess is a SettlementStatus of type Success
	SettlementStatusSuccess
	// SettlementStatusFailure is a SettlementStatus of type Failure
	SettlementStatusFailure
	// SettlementStatusMissing is a SettlementStatus of type Missing
	SettlementStatusMissing
	// SettlementStatusDuplicate is a SettlementStatus of type Duplicate
	SettlementStatusDuplicate
)

const _SettlementStatusName = "unknownstage1successfailuremissingduplicate"

var _SettlementStatusNames = []string{
	_SettlementStatusName[0:7],
	_SettlementStatusName[7:13],
	_SettlementStatusName[13:20],
	_SettlementStatusName[20:27],
	_SettlementStatusName[27:34],
	_SettlementStatusName[34:43],
}

// SettlementStatusNames returns a list of possible string values of SettlementStatus.
func SettlementStatusNames() []string {
	tmp := make([]string, len(_SettlementStatusNames))
	copy(tmp, _SettlementStatusNames)
	return tmp
}

var _SettlementStatusMap = map[SettlementStatus]string{
	0: _SettlementStatusName[0:7],
	1: _SettlementStatusName[7:13],
	2: _SettlementStatusName[13:20],
	3: _SettlementStatusName[20:27],
	4: _SettlementStatusName[27:34],
	5: _SettlementStatusName[34:43],
}

// String implements the Stringer interface.
func (x SettlementStatus) String() string {
	if str, ok := _SettlementStatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("SettlementStatus(%d)", x)
}

var _SettlementStatusValue = map[string]SettlementStatus{
	_SettlementStatusName[0:7]:   0,
	_SettlementStatusName[7:13]:  1,
	_SettlementStatusName[13:20]: 2,
	_SettlementStatusName[20:27]: 3,
	_SettlementStatusName[27:34]: 4,
	_SettlementStatusName[34:43]: 5,
}

// ParseSettlementStatus attempts to convert a string to a SettlementStatus
func ParseSettlementStatus(name string) (SettlementStatus, error) {
	if x, ok := _SettlementStatusValue[name]; ok {
		return x, nil
	}
	return SettlementStatus(0), fmt.Errorf("%s is not a valid SettlementStatus, try [%s]", name, strings.Join(_SettlementStatusNames, ", "))
}

// MarshalText implements the text marshaller method
func (x SettlementStatus) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *SettlementStatus) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseSettlementStatus(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
