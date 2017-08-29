package logutil

func FormatMsg(model string, method string, uid string, val string, msg string) string {
	logstr := model + " " + method + " " + uid + " " + val + " " + msg
	return logstr
}
