package control

func ifNewVariable(start, end int) {
	if distance := end - start; distance > 100 {
		println("too far")
	} else if distance > 50 {
		println("a bit far")
	} else if distance < 20 {
		println("too close")
	}
	//new variable distance is only visible in if block
	// println(distance)
}
