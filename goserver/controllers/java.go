package controllers

import (
	//"encoding/json"
	//"goserver/models"

	"log"

	"encoding/json"
	"reflect"

	"github.com/astaxie/beego"
	"github.com/ggwhite/go-hessian"
)

// Operations about Users
type JavaController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *JavaController) GetAll() {
	var addr string = "http://localhost:8090/HASDXMK/remoting/xmkBmpmanagementService"
	proxy, err := hessian.NewProxy(&hessian.ProxyConfig{
		Version: hessian.V1,
		URL:     addr,
	})
	if err != nil {
		panic(err)
	}
	//public List<Map<String, Object>> getCtAgency(String schema, String admdiv, String bdgyear)
	log.Println("-------------------------invoke-------------------------")
	data, err := proxy.Invoke2("getCtAgency", "FASPSD130000000", "130000000", "2020") //后台调用成功，反序列化时，这个库出现问题。
	log.Println("-------------------------end------------------")
	if err != nil {
		panic(err)
	}
	log.Println(data)
	//log.Println(reflect.TypeOf(data))
	//b, _ := data.([]map[string]interface{})

	log.Println(reflect.TypeOf(data).Kind()) //类型
	log.Println(reflect.ValueOf(data))
	s := reflect.ValueOf(data)
	for i := 0; i < s.Len(); i++ {
		m := s.Index(i)
		log.Println(m)
		log.Println(reflect.TypeOf(m).Kind()) //struct
		//log.Println(reflect.ValueOf(m))
		//str := json.Marshal(m.Interface())
		//log.Println("序列化后的字符串:", string(str))

		//将m中的数值的，从map改成int
		//log.Println(m2)
	}

	//异常：multiple-value json.Marshal() in single-value context 的解决办法：
	//https://stackoverflow.com/questions/45900911/multiple-value-json-marshal-in-single-value-context
	//str := json.Marshal(s)
	//log.Println(string(str))

	res, err := json.Marshal(data.(string))
	if err != nil {
		panic(err)
	}
	log.Println(res)
}
