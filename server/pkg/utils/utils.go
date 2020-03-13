package utils

import (
	"github.com/aofei/sandid"
)

func GenUUID() string {
	uid := sandid.New()
	//uid, _ := uuid.NewUUID()
	return uid.String()
}

func Respond(err bool, data *string) map[string]interface{} {
	return map[string]interface{}{"error": err, "data": data}
}
