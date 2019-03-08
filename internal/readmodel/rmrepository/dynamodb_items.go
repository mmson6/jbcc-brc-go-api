package rmrepository

import (
	"github.com/jbcc/brc-api/internal/models"
	"github.com/jbcc/brc-api/pkg/brcapiv1"
	"github.com/mitchellh/mapstructure"
)

type tableItem struct {
	Chapters     int         `json:"chapters,omitemptry"`
	Data         interface{} `json:"data,omitempty"` // bible record
	Group        string      `json:"group"`
	PartitionKey string      `json:"pk"` // PK
	SortKey      string      `json:"sk"` // SK
	UniqueKey    string      `json:"unique_key"`
	Verses       int         `json:"verses,omitempty"`
}

const (
	sortKeyUserRecord  = "USR_RECORD"
	sortKeyUserProfile = "USR_PROFILE"
)

////////////////////////////////////////////////////////////////////////////////
// User

func userRecordForItem(item tableItem) *models.UserRecord {
	var record brcapiv1.Record
	_ = mapstructure.Decode(item.Data, &record)

	model := models.UserRecord{
		ID:           item.PartitionKey,
		Group:        item.Group,
		Record:       record,
		ChapterCount: item.Chapters,
		DisplayName:  item.UniqueKey,
		VerseCount:   item.Verses,
	}

	return &model
}
