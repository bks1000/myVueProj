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
	//var addr string = "http://172.18.5.171:8090/HASDXMK/remoting/xmkBmpmanagementService"
	var addr string = "http://192.168.7.102:8190/HebFASPSD/remoting/authorizationWithApplicationService"
	proxy, err := hessian.NewProxy(&hessian.ProxyConfig{
		Version: hessian.V1,
		URL:     addr,
	})
	if err != nil {
		panic(err)
	}
	//public List<Map<String, Object>> getCtAgency(String schema, String admdiv, String bdgyear)
	log.Println("-------------------------invoke-------------------------")
	//data, err := proxy.Invoke2("getCtAgency", "FASPSD130000000", "130000000", "2020", 1) //后台调用成功，反序列化时，这个库出现问题。//分多个包发来的，就不行了，回头再试试。
	data, err := proxy.Invoke2("getApplications", "FASPSD130000000", "") //Response Head中含有Transfer-Encoding: chunked 的，可以。
	if err != nil {
		panic(err)
	}
	//log.Println(data)
	log.Println("-------------------------end------------------")
	//log.Println(reflect.TypeOf(data))
	//b, _ := data.([]map[string]interface{})

	//log.Println(reflect.TypeOf(data).Kind()) //类型
	//log.Println(reflect.ValueOf(data))
	s := reflect.ValueOf(data)
	//log.Println(s)
	for i := 0; i < s.Len(); i++ {
		m := s.Index(i)
		log.Println(m)

		//log.Println(reflect.TypeOf(m).Kind()) //struct
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

	res, err := json.Marshal(data.([]hessian.Any))
	if err != nil {
		panic(err)
	}
	//log.Println(string(res))
	u.Data["json"] = string(res) //要返回json这个是固定的。Data["json"]
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 string
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *JavaController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		var addr string = "http://localhost:8100/HASDXMK/remoting/xmkBmpmanagementService"
		proxy, err := hessian.NewProxy(&hessian.ProxyConfig{
			Version: hessian.V1,
			URL:     addr,
		})
		if err != nil {
			panic(err)
		}
		//public List<Map<String, Object>> getCtAgency(String schema, String admdiv, String bdgyear)
		log.Println("-------------------------invoke-------------------------")
		data, err := proxy.Invoke2("getCtAgency", "FASPSD130000000", "130000000", "2020", 1) //后台调用成功，反序列化时，这个库出现问题。//分多个包发来的，就不行了，回头再试试。
		if err != nil {
			panic(err)
		}
		//log.Println(data)
		log.Println("-------------------------end------------------")
		//log.Println(reflect.TypeOf(data))
		//b, _ := data.([]map[string]interface{})

		//log.Println(reflect.TypeOf(data).Kind()) //类型
		//log.Println(reflect.ValueOf(data))
		/*s := reflect.ValueOf(data)
		//log.Println(s)
		for i := 0; i < s.Len(); i++ {
			m := s.Index(i)
			log.Println(m)

			//log.Println(reflect.TypeOf(m).Kind()) //struct
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
		*/

		res, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		//log.Println(string(res))
		u.Data["json"] = string(res) //要返回json这个是固定的。Data["json"]
	}
	u.ServeJSON()
}
