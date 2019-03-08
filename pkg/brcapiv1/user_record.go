package brcapiv1

type UserRecord struct {
	ID     string `json:"id"`
	Record Record `json:"record"`
}

type Record struct {
	Genesis        []int `json:"genesis"`
	Exodus         []int `json:"exodus"`
	Leviticus      []int `json:"leviticus"`
	Numbers        []int `json:"numbers"`
	Deuteronomy    []int `json:"deuteronomy"`
	Joshua         []int `json:"joshua"`
	Judges         []int `json:"judges"`
	Ruth           []int `json:"ruth"`
	Samuel1        []int `json:"samuel_1"`
	Samuel2        []int `json:"samuel_2"`
	Kings1         []int `json:"kings_1"`
	Kings2         []int `json:"kings_2"`
	Chronicles1    []int `json:"chronicles_1"`
	Chronicles2    []int `json:"chronicles_2"`
	Ezra           []int `json:"ezra"`
	Nehemiah       []int `json:"nehemiah"`
	Esther         []int `json:"esther"`
	Job            []int `json:"job"`
	Psalm          []int `json:"psalm"`
	Proverbs       []int `json:"proverbs"`
	Ecclesiastes   []int `json:"ecclesiastes"`
	SongOfSolomon  []int `json:"songOfSolomon"`
	Isaiah         []int `json:"isaiah"`
	Jeremiah       []int `json:"jeremiah"`
	Lamentations   []int `json:"lamentations"`
	Ezekiel        []int `json:"ezekiel"`
	Daniel         []int `json:"daniel"`
	Hosea          []int `json:"hosea"`
	Joel           []int `json:"joel"`
	Amos           []int `json:"amos"`
	Obadiah        []int `json:"obadiah"`
	Jonah          []int `json:"jonah"`
	Micah          []int `json:"micah"`
	Nahum          []int `json:"nahum"`
	Habakkuk       []int `json:"habakkuk"`
	Zephaniah      []int `json:"zephaniah"`
	Haggai         []int `json:"haggai"`
	Zechariah      []int `json:"zechariah"`
	Malachi        []int `json:"malachi"`
	Matthew        []int `json:"matthew"`
	Mark           []int `json:"mark"`
	Luke           []int `json:"luke"`
	John           []int `json:"john"`
	Acts           []int `json:"acts"`
	Romans         []int `json:"romans"`
	Corinthians1   []int `json:"corinthians_1"`
	Corinthians2   []int `json:"corinthians_2"`
	Galatians      []int `json:"galatians"`
	Ephesians      []int `json:"ephesians"`
	Philippians    []int `json:"philippians"`
	Colossians     []int `json:"colossians"`
	Thessalonians1 []int `json:"thessalonians_1"`
	Thessalonians2 []int `json:"thessalonians_2"`
	Timothy1       []int `json:"timothy_1"`
	Timothy2       []int `json:"timothy_2"`
	Titus          []int `json:"titus"`
	Philemon       []int `json:"philemon"`
	Hebrews        []int `json:"hebrews"`
	James          []int `json:"james"`
	Peter1         []int `json:"peter_1"`
	Peter2         []int `json:"peter_2"`
	John1          []int `json:"john_1"`
	John2          []int `json:"john_2"`
	John3          []int `json:"john_3"`
	Jude           []int `json:"jude"`
	Revelation     []int `json:"revelation"`
}
