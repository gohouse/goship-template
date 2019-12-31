package helper

import "github.com/gohouse/gorose/v2"

//递归实现(返回树状结果得数据)
func GetCategoryTree(allCate []gorose.Data, p_name interface{}) []gorose.Data {
	var arr []gorose.Data
	for _, v := range allCate {
		if p_name == v["p_name"] {
			var child = v
			child["child"] = GetCategoryTree(allCate, v["resource_name"])
			arr = append(arr, child)
		}
	}
	return arr
}
