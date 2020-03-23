package room

import (
	"go_sugared/pkg/util"
	"go_sugared/routers/api"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 数据库相关操作
func connectRoomMgo(db, cl string) *mgo.Collection {
	collection := api.MgoSession.DB(db).C(cl)
	return collection
}

func (doc *Room) Insert() error {
	c := connectRoomMgo("hh", "room")
	//defer api.MgoSession.Close()
	doc.PackageNumber = util.MakePackageNumber()
	return c.Insert(doc)
}

func (doc *DeleteRoomReq) Delete() error {
	c := connectRoomMgo("hh", "room")
	//defer api.MgoSession.Close()
	return c.Insert(doc)
}

func FindOneRoomByPackageNumber(pkg string) (error, interface{}) {
	var result Room
	var err error
	c := connectRoomMgo("hh", "room")
	err = c.Find(bson.M{"packagenumber": pkg}).Select(bson.M{"_id": 0}).One(&result)
	if err != nil {
		return err, nil
	}
	return nil, result
}

func findOneBySelector() {

}

func findAllBySelector() {

}