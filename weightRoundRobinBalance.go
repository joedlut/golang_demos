package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
3. 算法说明
在 Nginx加权轮询算法 中，每个节点都有3个权重的变量
Weight : 配置的权重，根据配置文件初始化每个服务器节点的权重
currentWeight : 节点的当前权重，初始化时是配置的权重，随后会一直变更
effectiveWeight : 有效的权重，初始值为 weight ，通讯过程中发现节点异常，则 -1 ，之后再次选择本节点，调用成功一次则 +1 ，直到恢复到 weight。这个参数可以用于做降权。或者说是你的设置的权限修正。。

Nginx加权轮询算法 的逻辑实现
轮询所有节点，计算当前状态下所有的节点的 effectiveWeight 之和 作为 totalWeight；
更新每个节点的 currentWeight ， currentWeight = currentWeight + effectiveWeight; 选出所有节点 currentWeight 中最大的一个节点作为选中节点；
选择中的节点再次更新 currentWeight, currentWeight = currentWeight - totalWeight；
*/

type WeightRoundRobinBalance struct {
	curIndex int
	rss      []*WeightNode
}
type WeightNode struct {
	weight          int
	currentWeight   int
	effectiveWeight int
	addr            string
}

func (r *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("params at least 2")
	}
	addr := params[0]
	weight, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}
	node := &WeightNode{
		weight:          int(weight),
		effectiveWeight: int(weight),
		currentWeight:   int(weight),
		addr:            addr,
	}
	r.rss = append(r.rss, node)
	return nil
}

func (r *WeightRoundRobinBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}
	totalWeight := 0
	var maxWeightNode *WeightNode
	for key, node := range r.rss {
		totalWeight += node.effectiveWeight
		node.currentWeight += node.effectiveWeight

		if maxWeightNode == nil || maxWeightNode.currentWeight < node.currentWeight {
			maxWeightNode = node
			r.curIndex = key
		}
	}
	maxWeightNode.currentWeight -= totalWeight
	return maxWeightNode.addr
}

func main() {
	var wrb WeightRoundRobinBalance
	wrb.Add("172.0.0.1:8001", "4")
	wrb.Add("172.0.0.1:8002", "2")
	wrb.Add("172.0.0.1:8003", "1")
	for i := 0; i < 7; i++ {
		fmt.Printf("%s %d\n", wrb.Next(), wrb.curIndex)
	}
}
