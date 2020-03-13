package persistance

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type persistance struct{ Data *leveldb.DB }

func New() *persistance {
	db, _ := leveldb.OpenFile("leveldb", nil)
	return &persistance{Data: db}
}

func (p *persistance) GetLink(value string) (*string, bool) {
	db := p.Data
	defer db.Close()

	item, err := db.Get([]byte(value), nil)

	if err != nil {
		return nil, true
	}

	data := fmt.Sprintf("%s", item)
	return &data, false
}

func (p *persistance) PutLink(uid string, link string) *string {
	db := p.Data
	defer db.Close()

	err := db.Put([]byte(uid), []byte(link), nil)
	if err != nil {
		return nil
	}

	return &uid
}
