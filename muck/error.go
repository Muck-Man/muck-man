package muck

func report(e error) {
	println("(muck err!)", e.Error())
}
