// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными
// параметрами x,y и конструктором.

type Point structure {
	x float64
	y float64
}

func distance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x - p2.x, 2) + math.Pow(p1.y - p2.y, 2))
}