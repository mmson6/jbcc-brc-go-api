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

	record = fillUnderscoreBookError(record, item)
	record = transformNullRecordToEmptyArray(record)
	
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

func leaderboardForItems(items []tableItem) *models.Leaderboard {
	users := make([]models.User, 0, len(items))
	for _, item := range items {
		user := models.User{
			DisplayName: item.UniqueKey,
			VerseCount:  item.Verses,
		}
		users = append(users, user)
	}

	leaderboard := models.Leaderboard{
		Users: users,
	}

	return &leaderboard
}

////////////////////////////////////////////////////////////////////////////////
// HELPER FUNCTIONS

func fillUnderscoreBookError(record brcapiv1.Record, item tableItem) brcapiv1.Record {
	recordMap := item.Data.(map[string]interface{})

	if recordMap["timothy_1"] != nil {
		tim1 := recordMap["timothy_1"].([]interface{})
		record.Timothy1 = transformToIntArray(tim1)
	}
	
	if recordMap["samuel_1"] != nil {
		samuel1 := recordMap["samuel_1"].([]interface{})
		record.Samuel1 = transformToIntArray(samuel1)
	}
	
	if recordMap["samuel_2"] != nil {
		samuel2 := recordMap["samuel_2"].([]interface{})
		record.Samuel2 = transformToIntArray(samuel2)
	}
	
	if recordMap["kings_1"] != nil {
		kings1 := recordMap["kings_1"].([]interface{})
		record.Kings1 = transformToIntArray(kings1)
	}
	
	if recordMap["kings_2"] != nil {
		kings2 := recordMap["kings_2"].([]interface{})
		record.Kings2 = transformToIntArray(kings2)
	}
	
	if recordMap["chronicles_1"] != nil {
		chronicles1 := recordMap["chronicles_1"].([]interface{})
		record.Chronicles1 = transformToIntArray(chronicles1)
	}
	
	if recordMap["chronicles_2"] != nil {
		chronicles2 := recordMap["chronicles_2"].([]interface{})
		record.Chronicles2 = transformToIntArray(chronicles2)
	}
	
	if recordMap["corinthians_1"] != nil {
		corinthians1 := recordMap["corinthians_1"].([]interface{})
		record.Corinthians1 = transformToIntArray(corinthians1)
	}
	
	if recordMap["corinthians_2"] != nil {
		corinthians2 := recordMap["corinthians_2"].([]interface{})
		record.Corinthians2 = transformToIntArray(corinthians2)
	}
	
	if recordMap["thessalonians_1"] != nil {
		thessalonians1 := recordMap["thessalonians_1"].([]interface{})
		record.Thessalonians1 = transformToIntArray(thessalonians1)
	}
	
	if recordMap["thessalonians_2"] != nil {
		thessalonians2 := recordMap["thessalonians_2"].([]interface{})
		record.Thessalonians2 = transformToIntArray(thessalonians2)
	}
	
	if recordMap["timothy_1"] != nil {
		timothy1 := recordMap["timothy_1"].([]interface{})
		record.Timothy1 = transformToIntArray(timothy1)
	}
	
	if recordMap["timothy_2"] != nil {
		timothy2 := recordMap["timothy_2"].([]interface{})
		record.Timothy2 = transformToIntArray(timothy2)
	}
	
	if recordMap["peter_1"] != nil {
		peter1 := recordMap["peter_1"].([]interface{})
		record.Peter1 = transformToIntArray(peter1)
	}
	
	if recordMap["peter_2"] != nil {
		peter2 := recordMap["peter_2"].([]interface{})
		record.Peter2 = transformToIntArray(peter2)
	}
	
	if recordMap["john_1"] != nil {
		john1 := recordMap["john_1"].([]interface{})
		record.John1 = transformToIntArray(john1)
	}
	
	if recordMap["john_2"] != nil {
		john2 := recordMap["john_2"].([]interface{})
		record.John2 = transformToIntArray(john2)
	}
	
	if recordMap["john_3"] != nil {
		john3 := recordMap["john_3"].([]interface{})
		record.John3 = transformToIntArray(john3)
	}
	
	return record
}

func transformToIntArray(data []interface{}) []int {
	array := make([]int, 0, len(data))
	for _, item := range data {
		intValue := int(item.(float64))
		array = append(array, intValue)
	}
	return array
}

func transformNullRecordToEmptyArray(record brcapiv1.Record) brcapiv1.Record {
	if record.Genesis == nil {
		record.Genesis = []int{}
	}
	if record.Exodus == nil {
		record.Exodus = []int{}
	}
	if record.Leviticus == nil {
		record.Leviticus = []int{}
	}
	if record.Numbers == nil {
		record.Numbers = []int{}
	}
	if record.Deuteronomy == nil {
		record.Deuteronomy = []int{}
	}
	if record.Joshua == nil {
		record.Joshua = []int{}
	}
	if record.Judges == nil {
		record.Judges = []int{}
	}
	if record.Ruth == nil {
		record.Ruth = []int{}
	}
	if record.Samuel1 == nil {
		record.Samuel1 = []int{}
	}
	if record.Samuel2 == nil {
		record.Samuel2 = []int{}
	}
	if record.Kings1 == nil {
		record.Kings1 = []int{}
	}
	if record.Kings2 == nil {
		record.Kings2 = []int{}
	}
	if record.Chronicles1 == nil {
		record.Chronicles1 = []int{}
	}
	if record.Chronicles2 == nil {
		record.Chronicles2 = []int{}
	}
	if record.Ezra == nil {
		record.Ezra = []int{}
	}
	if record.Nehemiah == nil {
		record.Nehemiah = []int{}
	}
	if record.Esther == nil {
		record.Esther = []int{}
	}
	if record.Job == nil {
		record.Job = []int{}
	}
	if record.Psalm == nil {
		record.Psalm = []int{}
	}
	if record.Proverbs == nil {
		record.Proverbs = []int{}
	}
	if record.Ecclesiastes == nil {
		record.Ecclesiastes = []int{}
	}
	if record.SongOfSolomon == nil {
		record.SongOfSolomon = []int{}
	}
	if record.Isaiah == nil {
		record.Isaiah = []int{}
	}
	if record.Jeremiah == nil {
		record.Jeremiah = []int{}
	}
	if record.Lamentations == nil {
		record.Lamentations = []int{}
	}
	if record.Ezekiel == nil {
		record.Ezekiel = []int{}
	}
	if record.Daniel == nil {
		record.Daniel = []int{}
	}
	if record.Hosea == nil {
		record.Hosea = []int{}
	}
	if record.Joel == nil {
		record.Joel = []int{}
	}
	if record.Amos == nil {
		record.Amos = []int{}
	}
	if record.Obadiah == nil {
		record.Obadiah = []int{}
	}
	if record.Jonah == nil {
		record.Jonah = []int{}
	}
	if record.Micah == nil {
		record.Micah = []int{}
	}
	if record.Nahum == nil {
		record.Nahum = []int{}
	}
	if record.Habakkuk == nil {
		record.Habakkuk = []int{}
	}
	if record.Zephaniah == nil {
		record.Zephaniah = []int{}
	}
	if record.Haggai == nil {
		record.Haggai = []int{}
	}
	if record.Zechariah == nil {
		record.Zechariah = []int{}
	}
	if record.Malachi == nil {
		record.Malachi = []int{}
	}
	if record.Matthew == nil {
		record.Matthew = []int{}
	}
	if record.Mark == nil {
		record.Mark = []int{}
	}
	if record.Luke == nil {
		record.Luke = []int{}
	}
	if record.John == nil {
		record.John = []int{}
	}
	if record.Acts == nil {
		record.Acts = []int{}
	}
	if record.Romans == nil {
		record.Romans = []int{}
	}
	if record.Corinthians1 == nil {
		record.Corinthians1 = []int{}
	}
	if record.Corinthians2 == nil {
		record.Corinthians2 = []int{}
	}
	if record.Galatians == nil {
		record.Galatians = []int{}
	}
	if record.Ephesians == nil {
		record.Ephesians = []int{}
	}
	if record.Philippians == nil {
		record.Philippians = []int{}
	}
	if record.Colossians == nil {
		record.Colossians = []int{}
	}
	if record.Thessalonians1 == nil {
		record.Thessalonians1 = []int{}
	}
	if record.Thessalonians2 == nil {
		record.Thessalonians2 = []int{}
	}
	if record.Timothy1 == nil {
		record.Timothy1 = []int{}
	}
	if record.Timothy2 == nil {
		record.Timothy2 = []int{}
	}
	if record.Titus == nil {
		record.Titus = []int{}
	}
	if record.Philemon == nil {
		record.Philemon = []int{}
	}
	if record.Hebrews == nil {
		record.Hebrews = []int{}
	}
	if record.James == nil {
		record.James = []int{}
	}
	if record.Peter1 == nil {
		record.Peter1 = []int{}
	}
	if record.Peter2 == nil {
		record.Peter2 = []int{}
	}
	if record.John1 == nil {
		record.John1 = []int{}
	}
	if record.John2 == nil {
		record.John2 = []int{}
	}
	if record.John3 == nil {
		record.John3 = []int{}
	}
	if record.Jude == nil {
		record.Jude = []int{}
	}
	if record.Revelation == nil {
		record.Revelation = []int{}
	}
	return record
}
