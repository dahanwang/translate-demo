package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"translate-demo/models"
	"translate-demo/utils"
	"github.com/astaxie/beego"
)

// Operations about Users
type TranslateController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *TranslateController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *TranslateController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *TranslateController) Get() {
	//data := u.GetString("data")
	//hash := md5.Sum([]byte(data))
	//u.Ctx.WriteString("data="+data+",hash="+hex.EncodeToString(hash[:]))

	data := "Coronaviruses are a large family of viruses that cause disease in mammals and birds. Coronaviruses can cause illnesses that range from the common cold to much more severe illnesses like Middle East Respiratory Syndrome (MERS) and Severe Acute Respiratory Syndrome (SARS)."
	data = strings.Replace(data, " ", "", -1)
	//data := "Hello"
	transDataStr := translateByBaidu(data,"auto", "zh")
	fmt.Println("transdata = " + transDataStr)

	var transData trans_data
	if err := json.Unmarshal([]byte(transDataStr), &transData); err != nil {
		fmt.Println(err)
	}

	fmt.Println(transData)
	for i:=0; i < len(transData.Trans_result); i++ {
		result:=transData.Trans_result[i]
		u.Ctx.WriteString(result.Dst)
	}
}

type trans_data struct {
	From string `json:"from"`
	To  string `json:"to"`
	Trans_result []trans_result `json:"trans_result"`
}

type trans_result struct {
	Src string `json:"src"`
	Dst  string `json:"dst"`
}

var appid = "20200213000383367"
func translateByBaidu(q string,from string,to string) string {
	salt := strconv.Itoa(rand.Intn(10000))
	data := appid + q +salt+ "Z09vzLtis5ZjVDKWwhbX"
	signHash := md5.Sum([]byte(data))
	url := "http://api.fanyi.baidu.com/api/trans/vip/translate?"
	url +="q="+q
	url += "&from="+from
	url += "&to="+to
	url += "&appid=" + appid
	url += "&salt=" + salt
	url += "&sign=" + hex.EncodeToString(signHash[:])

	response := utils.HttpGet(url)

	return response
	//fmt.Println(response)
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *TranslateController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *TranslateController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

