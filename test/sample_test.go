package test

import (
	"gin-sam/comm"
	"testing"
	"time"
)

// 声明输入、输出
type TestData struct {
	in  int64
	out string
}

// 构造数据表格
var data = []TestData{
	{
		in:  1587545588,
		out: "2020-04-22 16:53:08",
	},
	{
		0, time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
	},
	{
		-2, time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"),
	},
}

// 执行测试
func TestMe(t *testing.T) {
	for _, d := range data {
		n := comm.FormatFromUnixTime(d.in)
		if n != d.out {
			t.Errorf("FormatFromUnixTime(%d) = %s;should be %s", d.in, n, d.out)
		}
	}
}
