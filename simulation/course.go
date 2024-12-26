package simulation

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"strconv"
	"time"
)

type courseBooks struct {
	BookId           string  `json:"bookId"`
	BookName         string  `json:"bookName"`
	BookType         int     `json:"bookType"`
	CategoryType     string  `json:"categoryType"`
	FileId           string  `json:"fileId"`
	GoodsId          int     `json:"goodsId"`
	GoodsImg         string  `json:"goodsImg"`
	GoodsPrice       float64 `json:"goodsPrice"`
	GoodsVerticalImg string  `json:"goodsVerticalImg"`
	IsGoods          int     `json:"isGoods"`
	IsShelf          int     `json:"isShelf"`
	MarketPrice      float64 `json:"marketPrice"`
	PackType         string  `json:"packType"`
	PcFileId         string  `json:"pcFileId"`
	SpecialFlag      int     `json:"specialFlag"`
	StageId          int     `json:"stageId"`
	StudyTraceFlag   int     `json:"studyTraceFlag"`
	SubjectId        int     `json:"subjectId"`
	ToMobile         int     `json:"toMobile"`
	ToPC             int     `json:"toPC"`
	WordsFlag        int     `json:"wordsFlag"`
}

/**
 * 课程返回信息 包含课程数据 | 返回结果
 */
type courseInfo struct {
	Data []struct {
		Books      []courseBooks `json:"books"`
		CourseCode string        `json:"courseCode"`
		CourseID   int           `json:"courseId"`
		CourseName string        `json:"courseName"`
		IsGoods    int           `json:"isGoods"`
		IsShelf    int           `json:"isShelf"`
		OrganID    int           `json:"organId"`
		SubjectID  int           `json:"subjectId"`
	} `json:"data"`
	Result Result `json:"result"`
}

// ProcessInfo 进度成绩信息
type ProcessInfo struct {
	Data []struct {
		BookID           string  `json:"bookId"`
		BookName         string  `json:"bookName"`
		CourseID         int     `json:"courseId"`
		GoodsImg         string  `json:"goodsImg"`
		GoodsVerticalImg string  `json:"goodsVerticalImg"`
		Percent          float64 `json:"percent"`
		Score            float64 `json:"score"`
		Seconds          int     `json:"seconds"`
		UID              int     `json:"uid"`
	} `json:"data"`
	Result Result `json:"result"`
}

type BookInfo struct {
	Data   []BookData `json:"data"`
	Result Result     `json:"result"`
}

// BookData 当前所拥有的教材数据
type BookData struct {
	AddTime          int64   `json:"addTime"`
	AutherName       string  `json:"autherName,omitempty"`
	BookId           string  `json:"bookId"`
	BookName         string  `json:"bookName"`
	BookType         int     `json:"bookType"`
	CategoryType     int     `json:"categoryType"`
	EndDate          int64   `json:"endDate"`
	FileId           string  `json:"fileId,omitempty"`
	GoodsId          int     `json:"goodsId"`
	GoodsImg         string  `json:"goodsImg"`
	GoodsPrice       float64 `json:"goodsPrice"`
	GoodsVerticalImg string  `json:"goodsVerticalImg"`
	IsGoods          int     `json:"isGoods"`
	IsShelf          int     `json:"isShelf"`
	PackType         int     `json:"packType"`
	Password         string  `json:"password,omitempty"`
	PcFileId         string  `json:"pcFileId,omitempty"`
	PcVersion        string  `json:"pcVersion,omitempty"`
	SpId             int     `json:"spId"`
	StageId          int     `json:"stageId"`
	StartDate        int64   `json:"startDate"`
	SubjectId        int     `json:"subjectId"`
	ToMobile         int     `json:"toMobile"`
	ToPC             int     `json:"toPC"`
	ValidDays        int     `json:"validDays"`
}

// TreeInfo 课程章节信息遍历
type TreeInfo struct {
	Data struct {
		Pages []struct {
			Score        int    `json:"score,omitempty"`
			Name         string `json:"name"`
			DisplayOrder int    `json:"displayOrder"`
			ID           string `json:"id"`
			PlatformID   int    `json:"platformId"`
			Type         int    `json:"type"`
			ParentID     string `json:"parentId"`
			TaskID       string `json:"taskId"`
			BookID       string `json:"bookId"`
		} `json:"pages"`
		Chapters []struct {
			WelFileName  string  `json:"welFileName"`
			AddTime      int64   `json:"addTime"`
			Level        int     `json:"level"`
			IsContent    int     `json:"isContent"`
			DisplayOrder int     `json:"displayOrder"`
			Weight       float64 `json:"weight"`
			Type         int     `json:"type"`
			ParentID     string  `json:"parentId"`
			ExtendParams string  `json:"extendParams"`
			IsFree       int     `json:"isFree"`
			ChapterType  int     `json:"chapterType"`
			IsTest       int     `json:"isTest"`
			MobileStatus int     `json:"mobileStatus"`
			Name         string  `json:"name"`
			XotPaperID   string  `json:"xotPaperId"`
			PcStatus     int     `json:"pcStatus"`
			ID           string  `json:"id"`
			PcFileID     string  `json:"pcFileId,omitempty"`
			PcVersion    string  `json:"pcVersion,omitempty"`
			FileID       string  `json:"fileId,omitempty"`
		} `json:"chapters"`
		Book struct {
			PcFileID  string `json:"pcFileId"`
			PcVersion string `json:"pcVersion"`
			FileID    string `json:"fileId"`
		} `json:"book"`
		Tasks []struct {
			Score        float64 `json:"score"`
			TaskType     string  `json:"taskType"`
			DisplayOrder int     `json:"displayOrder"`
			Weight       float64 `json:"weight"`
			ID           string  `json:"id"`
			Type         int     `json:"type"`
			ParentID     string  `json:"parentId"`
			BookID       string  `json:"bookId"`
		} `json:"tasks"`
	} `json:"data"`
	Result Result `json:"result"`
}

// 提交数据结构体 用于转换json
type submitData struct {
	BookId       string `json:"bookId"`
	ChapterId    string `json:"chapterId"`
	ExtendParams struct {
	} `json:"extendParams"`
	Percent int `json:"percent"`
	QstJson []struct {
	} `json:"qstJson"`
	Result    string `json:"result"`
	Score     int    `json:"score"`
	Seconds   int    `json:"seconds"`
	StudyDate int64  `json:"studyDate"`
	TaskId    string `json:"taskId"`
	TaskNo    string `json:"taskNo"`
	Uid       int    `json:"uid"`
}

type DoCourseResult struct {
	Result Result `json:"result"`
}

// GetCourse 查课 返回当前所在班级内的课程信息（不论有没有购买课程）
func GetCourse(course *courseInfo, session *grequests.Session) {
	ticket := GetServerticket(session, COURSEURL)

	header["Host"] = "course-api.ismartlearning.cn"
	resp, err := session.Get(COURSEURL+TICKETS+ticket,
		&grequests.RequestOptions{
			Headers: header,
		})
	if err != nil {
		fmt.Printf("getcourse获取课程 生成结果解析错误：error:%s", err)
	} else {
		err := json.Unmarshal(resp.Bytes(), &course)
		if err != nil {
			fmt.Printf("getCourse Json解析失败：%s\n", err)
		}
		//获取到了课程列表

	}
}

func GetBuybooks(book *BookInfo, session *grequests.Session) {

	ticket := GetServerticket(session, BUYBOOKSURL)

	header["Host"] = "book-api.ismartlearning.cn"
	resp, err := session.Get(BUYBOOKSURL+TICKETS+ticket,
		&grequests.RequestOptions{
			Headers: header,
		})

	if err != nil {
		fmt.Printf("getbuybooks获取 生成结果解析错误：error:%s", err)
		return
	} else {
		err := json.Unmarshal(resp.Bytes(), &book)
		if err != nil {
			fmt.Printf("getBook Json解析失败：%s\n", err)
		} else {
			fmt.Println("Book数据已获取成功")
		}
		//获取到了课程列表
	}
}

// 获取所有章节
func GetTree(courseBook *courseBooks, session *grequests.Session) {

	ticket := GetServerticket(session, BOOKTREE)

	header["Host"] = "book-api.ismartlearning.cn"
	resp, err := session.Post(BOOKTREE+TICKETS+ticket,
		&grequests.RequestOptions{
			Data: map[string]string{
				"bookType": strconv.Itoa(courseBook.BookType),
				"bookId":   courseBook.BookId,
			},
			Headers: header,
		})
	if err != nil {
		fmt.Printf("get_tree 提交发生异常：%s", err)
	} else {
		err1 := json.Unmarshal(resp.Bytes(), &bookTree)
		if err1 != nil {
			fmt.Printf("getBook Json解析失败：%s\n", err1)
		} else {
			fmt.Println("Book数据已获取成功")
		}

	}

}

// 模拟学习
func doCourse(bookid string, chapterid string, taskid string, displaynum int, tic string, session *grequests.Session) bool {
	//Chapters中 "isContent":1, "level":0, "parentId":"", 判断是否为章节名称节点

	//模拟提交数据
	submit := submitData{
		BookId:       bookid,
		ChapterId:    chapterid,
		ExtendParams: struct{}{},
		Percent:      100,
		QstJson:      []struct{}{},
		Result:       "",
		Score:        RandInt64(96, 100),
		Seconds:      RandInt64(1000, 2000),
		StudyDate:    time.Now().Unix() * 1000,
		TaskId:       taskid,
		TaskNo:       strconv.Itoa(displaynum),
		Uid:          student.Data.Uid,
	}
	//结构体转json
	value, err := json.Marshal(&submit)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	taskJson := "[" + string(value) + "]"
	times := GetTime()
	timeMd5 := GetMd5(times)
	utValue := GetSubmitUt(taskJson, timeMd5)
	header["Host"] = "study-api.ismartlearning.cn"

	resp, err := session.Post(SUBMITURL+tic,
		&grequests.RequestOptions{
			Data: map[string]string{
				"tasksJson": taskJson,
				"ut":        utValue,
			},
			Headers: header,
		})
	if err != nil {
		fmt.Printf("do_course 提交发生异常：%s", err)
	} else {
		//fmt.Printf("do_course 提交成功：%s",string(resp.Bytes()))
		var res DoCourseResult
		err = json.Unmarshal(resp.Bytes(), &res)
		if res.Result.Msg == "success" {
			return true
		}
	}
	return false
}

// 模拟学习
func doCourse2(bookid string, chapterid string, taskid string, displaynum int, score int, tic string, session *grequests.Session) bool {
	//Chapters中 "isContent":1, "level":0, "parentId":"", 判断是否为章节名称节点

	//模拟提交数据
	submit := submitData{
		BookId:    bookid,
		ChapterId: chapterid,
		Percent:   100,
		QstJson:   []struct{}{},
		Result:    "",
		Score:     RandInt64(96, 100), //TODO 分数不确定
		Seconds:   RandInt64(1000, 2000),
		StudyDate: time.Now().Unix() * 1000,
		TaskId:    taskid,
		TaskNo:    strconv.Itoa(displaynum),
		Uid:       student.Data.Uid,
	}
	//结构体转json
	value, err := json.Marshal(&submit)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	taskJson := "[" + string(value) + "]"
	times := GetTime()
	timeMd5 := GetMd5(times)
	utValue := GetSubmitUt(taskJson, timeMd5)
	header["Host"] = "study-api.ismartlearning.cn"

	resp, err := session.Post(SUBMITURL+tic,
		&grequests.RequestOptions{
			Data: map[string]string{
				"tasksJson": taskJson,
				"ut":        utValue,
			},
			Headers: header,
		})
	if err != nil {
		fmt.Printf("do_course2 提交发生异常：%s", err)
	} else {
		var res DoCourseResult
		err = json.Unmarshal(resp.Bytes(), &res)
		if res.Result.Msg == "success" {
			return true
		}
	}
	return false
}

// GetProcess 查询进度
func GetProcess(session *grequests.Session, courseId int, processInfo *ProcessInfo) {
	ticket := GetServerticket(session, DASHBOARDURL)

	header["Host"] = "dsb-api.ismartlearning.cn"
	resp, err := session.Post(DASHBOARDURL+TICKETS+ticket,
		&grequests.RequestOptions{
			Data: map[string]string{
				"courseId": strconv.Itoa(courseId),
			},
			Headers: header,
		})
	if err != nil {
		fmt.Printf("GetProcess 提交发生异常：%s", err)
	} else {
		err1 := json.Unmarshal(resp.Bytes(), &processInfo)
		if err1 != nil {
			fmt.Printf("GetProcess Json解析失败：%s\n", err1)
		} else {
			fmt.Println("GetProcess 获取成绩成功")
		}
	}
}
