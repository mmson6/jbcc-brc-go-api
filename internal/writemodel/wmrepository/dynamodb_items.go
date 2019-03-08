package wmrepository

import (
	"github.com/jbcc/brc-api/internal/models"
)

type tableItem struct {
	Chapters     int         `json:"chapters,omitemptry"`
	Data         interface{} `json:"data,omitempty"` // bible record
	Group        string      `json:"group,omitempty"`
	PartitionKey string      `json:"pk"` // PK
	SortKey      string      `json:"sk"` // SK
	UniqueKey    string      `json:"unique_key,omitempty"`
	Verses       int         `json:"verses,omitempty"`
}

const (
	sortKeyUserRecord  = "USR_RECORD"
	sortKeyUserProfile = "USR_PROFILE"
)

////////////////////////////////////////////////////////////////////////////////
// User

func itemForUserProfile(model models.UserProfile) tableItem {
	item := tableItem{
		Group:        model.Group,
		PartitionKey: model.DisplayName,
		SortKey:      sortKeyUserProfile,
		UniqueKey:    model.ID,
	}

	return item
}

func itemForUserRecord(model models.UserRecord) tableItem {
	item := tableItem{
		Chapters:     model.ChapterCount,
		Data:         model.Record,
		Group:        model.Group,
		PartitionKey: model.ID,
		SortKey:      sortKeyUserRecord,
		UniqueKey:    model.DisplayName,
		Verses:       model.VerseCount,
	}

	return item
}
