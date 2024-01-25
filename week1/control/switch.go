package control

//don't need break in switch conditions
func SwitchVal(arg int) {
	switch arg {
	case 0:
		println("初始化")
	case 1:
		println("运行中")
	default:
		println("未知状态")
	}
}

func SwitchExpr(arg int) {
	switch {
	case arg > 100:
		println("too far")
	case arg > 50:
		println("safe distance")
	case arg < 20:
		println("too close")
	}
}
