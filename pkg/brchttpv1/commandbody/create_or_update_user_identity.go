package commandbody

import (
	"bytes"
	"encoding/json"

	"github.com/jbcc/brc-api/pkg/brcapiv1"
)

type CreateOrUpdateUserIdentity struct {
	Data brcapiv1.UserIdentity `json:"data"`
}

func (body CreateOrUpdateUserIdentity) JSONBody() (*bytes.Buffer, error) {
	binary, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(binary)
	return buf, nil
}
