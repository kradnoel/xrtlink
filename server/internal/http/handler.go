package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	"github.com/kradnoel/xrtlink/internal/persistance"
	u "github.com/kradnoel/xrtlink/pkg/utils"
	"github.com/osamingo/indigo"
)

var g *indigo.Generator
var message string

type Link struct {
	Link string `json:"link"`
}

func init() {
	t := time.Unix(1257894000, 0) // 2009-11-10 23:00:00 UTC
	g = indigo.New(nil, indigo.StartTime(t))
	_, err := g.NextID()
	if err != nil {
		log.Fatalln(err)
	}

	/*co := cors.Options{
		A
	}*/
}

/*var ping2 = func(c *gin.Context) {
	c.JSON(200, gin.H{"alive": "true"})
}*/

//var getLink = func(c *gin.Context) {
var getLink = func(w http.ResponseWriter, r *http.Request) {
	//query := c.Param("id")
	//query := r.Header.Get("id")
	query := chi.URLParam(r, "id")

	fmt.Println(query)

	persistance := persistance.New()
	data, isEmpty := persistance.GetLink(query)

	//c.Header("Access-Control-Allow-Origin", "*")

	if isEmpty == true {
		message = "Link not found!!!"
		//c.JSON(200, u.Respond(true, &message))
		json.NewEncoder(w).Encode(u.Respond(true, &message))
		return
	}

	json.NewEncoder(w).Encode(u.Respond(false, data))
	//c.JSON(200, u.Respond(false, data))
}

//var postLink = func(c *gin.Context) {
var postLink = func(w http.ResponseWriter, r *http.Request) {

	//link := c.PostForm("link")

	link := &Link{}

	err := json.NewDecoder(r.Body).Decode(link)

	if err != nil {
		log.Panic(err)
	}

	//fmt.Println("Link => " + link.Link)
	//fmt.Print("content-type: " + c.ContentType())
	//fmt.Println(c.GetHeader("Origin"))
	//c.SetAccepted("application/x-www-form-urlencoded;charset=utf-8")

	//c.Header("Access-Control-Allow-Origin", "*")

	//if link == "" || govalidator.IsURL(link) == false {
	//	message = "Shortlink is Invalid!!!"
	//c.JSON(200, u.Respond(true, &message))
	//	json.NewEncoder(w).Encode(u.Respond(true, &message))
	//	return
	//}

	//persistance := persistance.New()
	//uid, _ := g.NextID()
	//data := persistance.PutLink(uid, link)

	//c.JSON(200, u.Respond(false, data))
	//json.NewEncoder(w).Encode(u.Respond(false, data))
	//message := "goo.gl/xyz"

	if link.Link == "" || govalidator.IsRequestURL(link.Link) == false {
		message = "Link is Invalid!!! Try again!!!"
		json.NewEncoder(w).Encode(u.Respond(true, &message))
		return
	}

	persistance := persistance.New()
	uid, _ := g.NextID()
	data := persistance.PutLink(uid, link.Link)
	json.NewEncoder(w).Encode(u.Respond(false, data))
}

var optionsLink = func(c *gin.Context) {
	fmt.Println(c.GetHeader("Origin"))
	fmt.Print("content-type: " + c.ContentType())
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Done()
}

/*var SeedLinks = func(c *gin.Context) {
	link1 := t.NewLinkWithData("https://github.com/kradnoel/cambiomz", "2GR7A8Ni2C8")
	link2 := t.NewLinkWithData("https://github.com/kradnoel/xrtlink", "2GR7BvjwTp4")
	link3 := t.NewLinkWithData("https://cambiomz.herokuapp.com", "2GREoFKMf3v")
	link4 := t.NewLinkWithData("https://kradnoel.com", "2GSBL1goENC")

	db, _ := leveldb.OpenFile("leveldb", nil)
	defer db.Close()
	db := t.NewDB().Data
	defer db.Close()

	err := db.Put([]byte(link1.UID), []byte(link1.Link), nil)
	err = db.Put([]byte(link2.UID), []byte(link2.Link), nil)
	if err != nil {
		log.Panic(err)
	}
	c.JSON(200, gin.H{"data": "data inserted sucessfully"})
}*/
