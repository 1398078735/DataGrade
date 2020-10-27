package main

import (
	"fmt"
	"github.com/gly-hub/excel_utils"
	"io/ioutil"
)

//修改处把答案和获取到的答案放入一维数组再进行对比
func main() {
	//GetData()//获取答案
	//fmt.Println(datayes)

	dataclass3 := GetAllFile3("E:/GoProjects/src/统计分数/class3")
	fmt.Println(dataclass3)
	//dataclass4 := GetAllFile4("E:/GoProjects/src/统计分数/class4")
	//var answers4 []string
	//
	//
	//rd4, _ := ioutil.ReadDir("E:/GoProjects/src/统计分数/class4")
	//for _, fi := range rd4 {
	//	add := 0
	//	for _, answers4 = range dataclass4 {
	//		for i := 0; i < len(answers4); i++ {
	//			if answers[i] == answers4[i] {
	//				add += 1
	//			}
	//		}
	//	}
	//	fmt.Println(fi.Name()+"的分数:", add)
	//}

}

func GetData(excelpath string) [][]string {
	//excelpath = "genesis_file.xlsx"
	sheet := "Sheet1"
	col := make([]string, 7)
	col[0] = "题号"
	col[1] = "题目"
	col[2] = "A"
	col[3] = "B"
	col[4] = "C"
	col[5] = "D"
	col[6] = "答案"

	col1 := make([]int, 1)
	col1[0] = 6
	data, err := excel_utils.ReadExcel(excelpath, sheet)
	if err != nil {
		panic(err)
	} else {
		data = data.ValueLoc(col)
		//fmt.Println(data.ValueIndex(col1))
	}
	var datayes [][]string
	datayes = data.ValueIndex(col1).Data
	//fmt.Println(datayes)
	return datayes
}

func GetStudentData3(excelpath string) [][]string {
	datayes := GetData("genesis_file.xlsx")
	sheet := "Sheet1"
	col := make([]string, 7)
	col[0] = "题号"
	col[1] = "题目"
	col[2] = "A"
	col[3] = "B"
	col[4] = "C"
	col[5] = "D"
	col[6] = "答案"

	col1 := make([]int, 1)
	col1[0] = 6

	data, err := excel_utils.ReadExcel(excelpath, sheet)
	if err != nil {
		panic(err)
	} else {
		data = data.ValueLoc(col)
	}
	var dataclass3 [][]string
	dataclass3 = data.ValueIndex(col1).Data
	fmt.Println(datayes)
	fmt.Println(dataclass3)
	rd3, _ := ioutil.ReadDir("E:/GoProjects/src/统计分数/class3")
	for _, fi := range rd3 {
		add := 0
		for j := 0; j < 39;j++ {
			for i := 0; i < len(datayes[j]); i++ {
				if datayes[0][i] == dataclass3[j][i] {
					add++
				}
			}
		}
		fmt.Println(fi.Name()+"的分数:", add)
	}
	//fmt.Println(dataclass3)
	return dataclass3
}

func GetStudentData4(excelpath string) [][]string {
	sheet := "Sheet1"
	col := make([]string, 7)
	col[0] = "题号"
	col[1] = "题目"
	col[2] = "A"
	col[3] = "B"
	col[4] = "C"
	col[5] = "D"
	col[6] = "答案"

	col1 := make([]int, 1)
	col1[0] = 6

	data, err := excel_utils.ReadExcel(excelpath, sheet)
	if err != nil {
		panic(err)
	} else {
		data = data.ValueLoc(col)
		//fmt.Println(data.ValueIndex(col1))
	}
	var dataclass4 [][]string
	dataclass4 = data.ValueIndex(col1).Data
	//fmt.Println(dataclass4)
	return dataclass4
}

func GetAllFile3(pathname string) [][]string {
	var dataclass3 [][]string
	rd, _ := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile3(pathname + fi.Name() + "\\")
		} else {
			excelpath := pathname + "/" + fi.Name()
			dataclass3 = GetStudentData3(excelpath)
		}
	}
	return dataclass3
}

func GetAllFile4(pathname string) [][]string {
	var dataclass4 [][]string
	rd, _ := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile4(pathname + fi.Name() + "\\")
		} else {
			excelpath := pathname + "/" + fi.Name()
			dataclass4 = GetStudentData4(excelpath)
		}
	}
	return dataclass4
}
