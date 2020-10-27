package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"strings"
)
//最终版本2.0
//使用github的360EntSecGroup-Skylar/excelize
var answers = make([]string, 0)
func main() {
	AddAnswers()

	//把两个班的名字放进一个一维数组里
	class3Names := make([]string, 0)
	class4Names := make([]string, 0)

	//获取xlsx文件的名字
	class3Names = getFilenames("./class3", class3Names,"class3/")
	class4Names = getFilenames("./class4", class4Names,"class4/")
	fmt.Println("三班成绩——————————————————————————————————————————————")
	GetGrade(class3Names,"class3/")
	fmt.Println("四班成绩——————————————————————————————————————————————")
	GetGrade(class4Names,"class4/")
}

//读取标准答案并且把标准答案追加到answers的一维数组中
func AddAnswers() {
	file, err := excelize.OpenFile("genesis_file.xlsx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	rows, err := file.GetRows("Sheet1")
	for i := 1; i < len(rows); i++ {
		answers = append(answers, rows[i][6])//获取到第7列答案那列并追加到answers数组中
	}

}

//获取到文件名字
func getFilenames(folder string, array []string,classname string) []string {
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		array = append(array, classname+file.Name())
	}
	return array
}

//得到三班和四班所有人的成绩并进行一一对比
func GetGrade(array []string,prefix string)  {
	for _, filename := range array {
		Answers := make([]string, 0)
		grade := 0
		gradeNil:= 0
		file, err := excelize.OpenFile(filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		Rows, err := file.GetRows("Sheet1")//读取文件中表名为sheet1的内容

		for i := 1; i < len(answers); i++ {//读取到答案列
			Answers = append(Answers, Rows[i][6])
		}

		for i := 0; i < len(Answers); i++ {
			if answers[i] == strings.ToTitle(Answers[i]) {//增加一个大小写转换
				grade++
			}
			if Answers[i] == ""{
				gradeNil++
			}
		}
		name := strings.TrimPrefix(filename, prefix)
		name = strings.TrimSuffix(name, ".xlsx")
		fmt.Printf("%s\t 实际答题%d题,\t 成绩为%d分\n",name,160-gradeNil,grade)

	}
}
