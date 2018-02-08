# gosegment

golang 版中文分词包, inspired from 盘古分词

Fork from [chuanyi/gosegment](https://github.com/chuanyi/gosegment)

## Install
```bash
go get -u github.com/WindomZ/gosegment
```

## Usage
```
seg := gosegment.NewSegment()
words := seg.DoSegment(`主要功能: 中英文分词，未登录词识别,多元歧义自动识别,全角字符识别能力`)
for cur := words.Front(); cur != nil; cur = cur.Next() {
    w := cur.Value.(*dict.WordInfo)
    // fmt.Println(w.Word, "(", w.Position, ",", w.Rank, ")")
    ......
}
```
