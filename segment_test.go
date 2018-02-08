package gosegment

import (
	"testing"

	"github.com/WindomZ/gosegment/dict"
	"github.com/WindomZ/testify/assert"
)

var seg *Segment

func init() {
	seg = NewSegment()
}

func TestSegment_DoSegment(t *testing.T) {
	cases := map[string]int{
		"盘古":     5,
		"分词":     5,
		"eaglet": 5,
		"字典":     5,
		"主要功能":   5,
		"词":      1,
		"90":     1,
		"300":    1,
		"600":    1,
		"1.8":    1,
	}
	ret := seg.DoSegment(`盘古分词 简介: 盘古分词 是由eaglet 开发的一款基于字典的中英文分词组件
主要功能: 中英文分词，未登录词识别,多元歧义自动识别,全角字符识别能力
主要性能指标:
分词准确度:90%以上
处理速度: 300-600KBytes/s Core Duo 1.8GHz
用于测试的句子:
长春市长春节致词
长春市长春药店
IＢM的技术和服务都不错
张三在一月份工作会议上说的确实在理
于北京时间5月10日举行运动会
我的和服务必在明天做好`)
	count := 0
	for cur := ret.Front(); cur != nil; cur = cur.Next() {
		w := cur.Value.(*dict.WordInfo)
		if rank, ok := cases[w.Word]; ok {
			assert.Equal(t, rank, w.Rank)
			count++
		}
	}
	assert.Equal(t, 15, count)
}
