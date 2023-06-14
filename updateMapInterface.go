package main

import "fmt"

//用于更新 map[string]interface{}的值
// m[keyName] 必须是一个map[string]interface{}   if keyName == key  map[keyName][key1]=value1

//m["intf1"]   map[string]interface{}
//m["intf1"]["lan"]  map[string]string
//m["person"]  map[string][string]

//m["book"] "golang"

//updateMapStringInterface(m,"lan","go","golang")
//  intf1 map[lan:map[html:html5]] ====>  intf1 map[lan:map[html:html5 go:golang]]

// updateMapStringInterface(m,"person","thirdnname","joedlut")
// person map[firstname:han secondname:guodong] ====> person map[firstname:han secondname:guodong thirdname:joedlut]
var flag bool = false

func updateMapStringInterface(m map[string]interface{}, key, key1, value1 string) error {
	fmt.Println("initial", m)
	for k, v := range m {
		//fmt.Println(k, "----", v)
		vMap, ok := v.(map[string]interface{})
		if ok {
			updateMapStringInterface(vMap, key, key1, value1)
		}
		//vMap1, ok := v.(map[string]string)
		vMap1, ok := v.(map[string]interface{})
		if k == key {

			fmt.Println("+++++", vMap1)
			vMap1[key1] = value1
			fmt.Println("+++++", vMap1)
			v1 := interface{}(vMap1)
			m[k] = v1
			flag = true
		}
	}
	fmt.Println("final", m)
	return nil
}

func main() {
	m := make(map[string]interface{})
	//m1 := make(map[string]string)
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})
	//m3 := make(map[string]string)
	m3 := make(map[string]interface{})
	m3["html"] = "html5"
	m2["lan"] = m3
	m1["firstname"] = "han"
	m1["secondname"] = "guodong"
	m["person"] = m1
	m["book"] = "golang"
	m["intf1"] = m2

	//fmt.Printf("%T\n", m["person"])
	//fmt.Printf("%T\n", m["intf1"])

	//updateMapStringInterface(m, "person", "test2", "test3")

	//updateMapStringInterface(m,"lan","go","golang")
	//  intf1 map[lan:map[html:html5]] ====>  intf1 map[lan:map[html:html5 go:golang]]

	fmt.Println(m)
	updateMapStringInterface(m, "lan", "go", "golang")

	//updateMapStringInterface(m,"person","thirdnname","joedlut")
	// person map[firstname:han secondname:guodong] ====> person map[firstname:han secondname:guodong thirdname:joedlut]
	//updateMapStringInterface(m, "person", "thirdnname", "joedlut")
	//fmt.Println(flag)
	if !flag {
		fmt.Println("更新map失败")
	} else {
		fmt.Println("更新map成功")
	}
	fmt.Println(m)
}
