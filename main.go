package main


import (
	_ "image/jpeg"
	"github.com/widuu/goini"
	"log"
)

type FaceDetctor struct {
	cascadeFile  string
	minSize      int
	maxSize      int
	shiftFactor  float64
	scaleFactor  float64
	iouThreshold float64
}

//保存要识别的头像库
var srcface string;
//头像库文件
var facefinder ="./facefinder";
//摄像头设备
var camdev="/dev/video0";


var threshold=3000;//像素对比参数



/*判断是否在家*/
func checkHome(){


}


func main() {
	if(srcface==""){
		log.Printf("srcface error\r\n")
		return ;
	}
	//启动录像
//	cam.StartCapture(camdev,25);
}

/*初始化*/
func init(){
	//读取配置文件
	conf := goini.SetConfig("./conf.ini")
	confarr := conf.ReadList()
	for index := 0; index < len(confarr); index++ {
		confmap := confarr[index]
		if _, ok := confmap["config"]; ok {
			//摄像头
			if _, ok := confmap["config"]["camdev"]; ok {
				camdev = confmap["config"]["camdev"]
			}
		}
	}

}



