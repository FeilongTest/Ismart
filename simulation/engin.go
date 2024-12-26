package simulation

import (
	"fmt"
	"github.com/levigross/grequests"
	"ismartTest/progressbar"
	"log"
	"time"
)

func Run(username string, password string, courseName string, result *map[string]string) {
	//初始化协议头
	header["Host"] = "sso.ismartlearning.cn"
	header["user-agent"] = "Android-Ismart-Moblie 2.4.3"
	//生成票据必须要X-Requested-With参数
	header["X-Requested-With"] = "android"
	//生成serviceTicket必须有content-type参数 否则会返回空
	header["content-type"] = "application/x-www-form-urlencoded"

	var orderProcess = make(map[string]string)

	session := grequests.NewSession(nil)
	var loginResult string
	Login(username, password, session, &loginResult)

	if CompareCompare(loginResult, LOGININFO) {
		//登录成功 延迟1秒
		time.Sleep(1 * time.Second)
		//var book BookInfo
		var couse courseInfo
		GetCourse(&couse, session)

		//GetBuybooks(&book,session)
		var bookList []courseBooks
		var i, j int

		for i = 0; i < len(couse.Data); i++ {
			for j = 0; j < len(couse.Data[i].Books); j++ {
				//if CompareCompare((couse.Data[i].Books[j].BookName),(courseName)){
				if CompareCompare(DealString(couse.Data[i].Books[j].BookName), DealString(courseName)) {

					fmt.Println("待刷课程的Bookid存在于已获取的Book列表中，可继续执行操作")
					bookList = append(bookList, couse.Data[i].Books[j])
					//break I //找到对应课程跳出循环
				}
			}
		}

		if len(bookList) == 0 {
			log.Println("没有找到待刷课程！")
			orderProcess["status"] = "3"
			orderProcess["process"] = "没有找到待刷课程，请检查后联系管理员！"
			*result = orderProcess
			return
		}

		for count, books := range bookList {
			couseBook := books
			GetTree(&couseBook, session) //成功获取到结构树
			fmt.Println("待刷课程列表已初始化，可继续执行操作")

			isFirstSubmit := true
			var b progressbar.Bar
			var isSuccessful bool
			if bookTree.Data.Tasks != nil {
				b.NewOption(0, int64(len(bookTree.Data.Tasks)-1))
				for count, task := range bookTree.Data.Tasks {
					//获取所有任务信息 这里暂时为单任务模式
					b.Play(int64(count))
					if isFirstSubmit {
						submitTicket := GetServerticket(session, SUBMITURL)
						isSuccessful = doCourse(task.BookID, task.ParentID, task.ID, task.DisplayOrder, TICKETS+submitTicket, session)
						isFirstSubmit = false
					} else {
						isSuccessful = doCourse(task.BookID, task.ParentID, task.ID, task.DisplayOrder, "", session)
					}
					if !isSuccessful {
						fmt.Println("提交发生异常 退出循环")
						break
					}
				}
				b.Finish()
			} else { //Task 节点为空则为另一种订单 需要采用pages节点信息提交
				b.NewOption(0, int64(len(bookTree.Data.Pages)-1))
				for count, pages := range bookTree.Data.Pages {
					//获取所有任务信息 这里暂时为单任务模式
					b.Play(int64(count))
					if isFirstSubmit {
						submitTicket := GetServerticket(session, SUBMITURL)
						isSuccessful = doCourse2(couseBook.BookId, pages.ParentID, pages.ID, pages.DisplayOrder, pages.Score, TICKETS+submitTicket, session)
						isFirstSubmit = false
					} else {
						isSuccessful = doCourse2(couseBook.BookId, pages.ParentID, pages.ID, pages.DisplayOrder, pages.Score, "", session)
					}
					if !isSuccessful {
						fmt.Println("提交发生异常 退出循环")
						break
					}
				}
				b.Finish()
			}

			//课程任务结束 查询课程进度、成绩等信息
			var processInfo ProcessInfo
			GetProcess(session, couse.Data[count].CourseID, &processInfo)
			for _, process := range processInfo.Data {
				if CompareCompare(DealString(process.BookName), DealString(courseName)) {
					orderProcess["status"] = "2"
					orderProcess["process"] = fmt.Sprintf("成绩%.2f 进度%.2f%% 已完成", process.Score, process.Percent)
					break
				}
			}
		}

	} else { //登录失败
		orderProcess["status"] = "3"
		orderProcess["process"] = loginResult
	}

	*result = orderProcess

}
