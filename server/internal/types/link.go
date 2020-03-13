package types

import (
	"log"

	"github.com/asaskevich/govalidator"
)

type link struct {
	ID   uint   `json:"id"`
	Link string `json:"link"`
	UID  string `json:"unique_id"`
}

/*func NewLink() link {
	return link{}
}*/

func NewLinkWithData(l string, uid string) link {
	validate := govalidator.IsURL(l)

	if validate == true {
		//uid := u.GenUUID()
		return link{Link: l, UID: uid}
	}
	log.Fatal("link is not a uri!!!")
	return link{}
}

func (l *link) String() string {
	str := "[Link] = " + l.Link + ", [UID] = " + l.UID
	return str
}
