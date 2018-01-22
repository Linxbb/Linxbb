//log.SetFlags(log.Llongfile)
//使log输出加上行数

import(
  "github.com/gocraft/dbr"
	"log"
	"os"
	"io"
)

var Warning,Error * log.Logger

func init() {
	errFile,err:=os.OpenFile("errors.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("打开日志文件失败：",err)
	}
  Info = log.New(os.Stdout,"Info:",log.Ldate | log.Ltime | log.Lshortfile)
  Error = log.New(io.MultiWriter(errFile),"Error:", log.Ltime | log.Lshortfile)
  输出到控制台和文件
  //Error = log.New(io.MultiWriter(os.Stderr,errFile),"Error:",log.Ldate | log.Ltime | log.Lshortfile)
}

func main(){
  //按格式输出
  Info.Println("。。。。。")
  Error.Println("。。。。。。")
}
