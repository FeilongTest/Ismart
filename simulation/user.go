package simulation

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"strconv"
	"strings"
)

type loginInfo struct {
	Data   loginInfoData `json:"data"`
	Result Result        `json:"result"`
}

type loginInfoData struct {
	Tgt      string `json:"tgt"`
	Uid      int    `json:"uid"`
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Realname string `json:"realname"`
}

// Result 返回结果结构体 可通用给所有json返回结果。
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ticketData struct {
	ServerTicket string `json:"serverTicket"`
}

type ticketInfo struct {
	Data   ticketData `json:"data"`
	Result Result     `json:"result"`
}

type studentInfo struct {
	Data   StudentData `json:"data"`
	Result Result      `json:"result"`
}

// StudentData 登录获取学生信息 所有json键
type StudentData struct {
	BirthDay     int    `json:"birthDay"`
	BirthMonth   int    `json:"birthMonth"`
	BirthYear    int    `json:"birthYear"`
	ClassCode    string `json:"classCode"`
	ClassId      int    `json:"classId"`
	ClassName    string `json:"className"`
	CreateTime   int64  `json:"createTime"`
	Email        string `json:"email"`
	FirstLogin   bool   `json:"firstLogin"`
	Gender       int    `json:"gender"`
	HeadPhoto    string `json:"headPhoto"`
	Mobile       string `json:"mobile"`
	OrganId      int    `json:"organId"`
	OrganName    string `json:"organName"`
	QqOpenid     string `json:"qqOpenid"`
	Realname     string `json:"realname"`
	Reserved1    string `json:"reserved1"`
	SchoolType   int    `json:"schoolType"`
	Status       int    `json:"status"`
	Uid          int    `json:"uid"`
	UserUuid     string `json:"userUuid"`
	Username     string `json:"username"`
	Utype        int    `json:"utype"`
	Uuid         string `json:"uuid"`
	WeiboOpenid  string `json:"weiboOpenid"`
	WeixinOpenid string `json:"weixinOpenid"`
}

func Login(user string, pwd string, session *grequests.Session, result *string) {

	user = strings.TrimSpace(user)
	pwd = strings.TrimSpace(pwd)
	times := GetTime()
	timeMd5 := GetMd5(times)
	pwdMd5 := GetMd5(pwd)
	finalPwdMd5 := GetMd5(pwdMd5 + "fa&s*l%$k!fq$k!ld@fjlk")
	utValue := GetUt(user, finalPwdMd5, timeMd5)

	resp, err := session.Post(LOGINURL,
		&grequests.RequestOptions{
			Data: map[string]string{
				"password": finalPwdMd5,
				"username": user,
				"ut":       utValue,
			},
			Headers: header,
		})
	if err != nil {
		fmt.Printf("请求错误:%s\n", err)
	} else {
		var logininfo loginInfo
		err1 := json.Unmarshal(resp.Bytes(), &logininfo)
		if err1 != nil {
			fmt.Println("loginInfo 账号登陆结果解析错误：error:", err1)
			return
		}

		//logininfo.Result.Msg = 生成Tgt成功 logininfo.Result.Code = -26
		if logininfo.Result.Msg == "生成Tgt成功" {
			//调用生成serverTicket需要一个cookies
			header["uid"] = strconv.Itoa(logininfo.Data.Uid)
			header["Cookie"] = "CASTGC=" + logininfo.Data.Tgt
			serviceTicket := GetServerticket(session, CURRENTUSER)
			//登陆最后一步 获取用户信息
			finishLogin(session, serviceTicket)
			//fmt.Println(serverticket)
			*result = LOGININFO

		} else {
			*result = logininfo.Result.Msg
		}

	}

}

func GetServerticket(session *grequests.Session, dataUrl string) string {
	resp, err := session.Post(SERVERTICKETURL,
		&grequests.RequestOptions{
			Data: map[string]string{
				"service": dataUrl,
			},
			Headers: header,
		})
	if err != nil {
		return fmt.Sprintf("获取serverTicket出错：%s", err)
	} else {
		//获取ticket成功
		var ticket ticketInfo
		err1 := json.Unmarshal(resp.Bytes(), &ticket)
		if err1 != nil {
			fmt.Printf("ticketInfo 生成结果解析错误：error:%s\n", err1)
		}
		return ticket.Data.ServerTicket
	}
}

func finishLogin(session *grequests.Session, ticket string) {
	header["Host"] = "school.ismartlearning.cn"
	resp, err := session.Post(LOGINTICKET+ticket,
		&grequests.RequestOptions{
			Data: map[string]string{
				"status": "0",
			},
			Headers: header,
		})
	if err != nil {
		fmt.Printf("获取UserTicket出错：%s", err)
	} else {
		//用户信息获取到json
		err1 := json.Unmarshal(resp.Bytes(), &student)
		if err1 != nil {
			fmt.Println("studentInfo 生成结果解析错误：error:", err1)
			return
		}
		fmt.Printf("用户登录成功!\ntgt为：%s \nUID为：%d\n", header["Cookie"], student.Data.Uid)

		//到这里完全登陆完成
	}
}
