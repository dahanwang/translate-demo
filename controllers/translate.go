package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"strconv"
	"strings"
	"translate-demo/models"
	"translate-demo/utils"
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
	translate(u)
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
// http://localhost:8080/odin/translate?to=zh&data=xx
func (u *TranslateController) Get() {
	translate(u)
}

func translate(u *TranslateController) {
	data := u.GetString("data")
	tolg := u.GetString("to") // 目标语言

	if tolg == "" {
		tolg = "en"
	}
	hash := md5.Sum([]byte(data))
	md5 := hex.EncodeToString(hash[:])

	cacheTranslate := models.HGetTranslate(md5, tolg)
	if cacheTranslate != "" {
		writeStr := "原文字：" + data +"\n\n"+ "目标语种："+tolg+"\n\n" + "翻译文字：" +cacheTranslate
		u.Ctx.WriteString("cached.....\n\n" + writeStr)
		return
	}

	// translate
	replaceData := strings.Replace(data, " ", "", -1)
	transDataStr := translateByBaidu(replaceData,"auto", tolg)
	fmt.Println("transdata = " + transDataStr)

	var transData trans_data
	if err := json.Unmarshal([]byte(transDataStr), &transData); err != nil {
		fmt.Println(err)
	}

	fmt.Println(transData)

	if len(transData.Trans_result) >1 {
		fmt.Println("transdata result len > 1")
	}

	result:=transData.Trans_result[0]
	writeStr := "原文字：" + data +"\n\n"+ "原语种："+transData.From+"\n\n"+ "目标语种："+tolg+"\n\n" + "翻译文字：" +result.Dst
	u.Ctx.WriteString(writeStr)

	/*
		for i:=0; i < len(transData.Trans_result); i++ {
			result:=transData.Trans_result[i]

			writeStr := "原文字：" + data +"\n\n"+ "原语种："+transData.From+"\n\n" + "翻译文字：" +result.Dst
			u.Ctx.WriteString(writeStr)
		}
	*/

	// cache
	models.HSetTranslat(md5,tolg, result.Dst)
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
	url +="q=" + q
	url += "&from=" + from
	url += "&to=" + to
	url += "&appid=" + appid
	url += "&salt=" + salt
	url += "&sign=" + hex.EncodeToString(signHash[:])

	response := utils.HttpGet(url)

	return response
}

