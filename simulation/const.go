package simulation

const (
	//封包URL
	LOGINURL        = "https://sso.ismartlearning.cn/v2/tickets-v2"
	CURRENTUSER     = "https://school.ismartlearning.cn/client/user/currentuser"
	COURSEURL       = "https://course-api.ismartlearning.cn/client/student/course/list"
	SERVERTICKETURL = "https://sso.ismartlearning.cn/v1/serviceTicket"
	BOOKTREE        = "https://book-api.ismartlearning.cn/client/books/tree"
	SUBMITURL       = "https://study-api.ismartlearning.cn/client/task/score/submit"
	BUYBOOKSURL     = "https://book-api.ismartlearning.cn/client/books/buy-book"
	DASHBOARDURL    = "https://dsb-api.ismartlearning.cn/client/dashbordv2-ismart/student/course/books/info"

	//封包需要生成ticket的URL
	LOGINTICKET = "https://school.ismartlearning.cn/client/user/currentuser?ticket="
	TICKETS     = "?ticket="

	//程序判断标志符号
	LOGININFO = "login success"
)

var (
	header   = make(map[string]string)
	student  studentInfo //全局用来存储学生信息
	bookTree TreeInfo    //全局用来存储课程章节信息
)

// 这里用来存储测试环境控制变量
var (
	Debug bool = true //测试环境开关
)
