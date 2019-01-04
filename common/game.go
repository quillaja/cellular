package common

/*
Things to specify
Window:
	title (initial)
	title (updated every 1 sec)
		as func(params) string
	width, height
	background color (or func that returns a bg color)

Performance/hit detection:
	number of partitions [1,n]

	Entities need for logic:
	func to call to update object position/state

	func to call to determine if the object should be put into a partition
		as func(part *Partition) bool
	func to call when the object is put into a partition
		as func(part *Partition)

	func to call to determine if collision happened
		as func(a, b *Point) bool
	func to call to perform action when collision is detected
		as func(a, b *Point)

Graphics/Drawing

	Entities need for drawing:
	func to call to prepare visual aka Draw()
		as func()
	func to return visual that has pixel `Draw(Target)`. generally IMDraw
		as func() imdraw.IMDraw

*/
