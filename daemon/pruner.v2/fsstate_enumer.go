// Code generated by "enumer -type=FSState -json"; DO NOT EDIT.

//
package pruner

import (
	"encoding/json"
	"fmt"
)

const _FSStateName = "FSStateInitializedFSStatePlanningFSStatePlanErrFSStateExecutingFSStateExecuteErrFSStateExecuteSuccess"

var _FSStateIndex = [...]uint8{0, 18, 33, 47, 63, 80, 101}

func (i FSState) String() string {
	if i < 0 || i >= FSState(len(_FSStateIndex)-1) {
		return fmt.Sprintf("FSState(%d)", i)
	}
	return _FSStateName[_FSStateIndex[i]:_FSStateIndex[i+1]]
}

var _FSStateValues = []FSState{0, 1, 2, 3, 4, 5}

var _FSStateNameToValueMap = map[string]FSState{
	_FSStateName[0:18]:   0,
	_FSStateName[18:33]:  1,
	_FSStateName[33:47]:  2,
	_FSStateName[47:63]:  3,
	_FSStateName[63:80]:  4,
	_FSStateName[80:101]: 5,
}

// FSStateString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FSStateString(s string) (FSState, error) {
	if val, ok := _FSStateNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to FSState values", s)
}

// FSStateValues returns all values of the enum
func FSStateValues() []FSState {
	return _FSStateValues
}

// IsAFSState returns "true" if the value is listed in the enum definition. "false" otherwise
func (i FSState) IsAFSState() bool {
	for _, v := range _FSStateValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for FSState
func (i FSState) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for FSState
func (i *FSState) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("FSState should be a string, got %s", data)
	}

	var err error
	*i, err = FSStateString(s)
	return err
}
