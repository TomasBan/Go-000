# 学习笔记

## 1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

代码如下：

main方法代表最外层的调用方，service代表调用链路中的某一个段。按照逐层上报，只处理一次的思路来思考。

```
func main() {
	err := service()
	if err != nil {
		fmt.Printf("main:%+v\n",err)
	}
}


func service() error {
	// 业务代码...
	return dao()
}

func dao() error {
	// db逻辑...
	return xerror.Wrapf(sql.ErrNoRows,"no rows")
}
```