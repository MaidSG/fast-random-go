package randomDataModule


import (
	"math/rand"
	"strconv"
)


// 创造随机的加减法题目
func GetRandomAdditonOrSubtraction() string {
	a := getRandomNumber()
    b := getRandomNumber()
    return generateExpression(a, b)
}


// 获取随机数
func getRandomNumber() int {
	return rand.Intn(100) + 10
}

// 生成表达式
func generateExpression(a,b int) string {
	if rand.Intn(2) == 0 {
		return strconv.Itoa(a) + "+" + strconv.Itoa(b)
	} else {
		return strconv.Itoa(a) + "-" + strconv.Itoa(b)
	}
}


