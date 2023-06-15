package main

import (
	"errors"
	"fmt"
)

type RoundRobinBalance struct {
	curIndex int
	rss      []string
}

func (r *RoundRobinBalance) Add(params string) error {
	if params == "" {
		return errors.New("please input one server")
	}
	r.rss = append(r.rss, params)
	return nil
}
func (r *RoundRobinBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}
	lens := len(r.rss)
	if r.curIndex > lens {
		r.curIndex = 0
	}
	curAdd := r.rss[r.curIndex]
	//类似环形队列
	r.curIndex = (r.curIndex + 1) % lens
	return curAdd
}

func main() {
	var r RoundRobinBalance
	r.Add("127.0.0.1:80")
	r.Add("127.0.0.1:81")
	r.Add("127.0.0.1:82")
	r.Add("127.0.0.1:83")
	r.Add("127.0.0.1:84")
	r.Add("127.0.0.1:85")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d -------> %s\n", i, r.Next())
	}
}
