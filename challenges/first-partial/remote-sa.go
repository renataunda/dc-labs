package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"math"
	"strings"
)

type Point struct {
	X, Y float64
}

type Line struct {
	A, B, C float64
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//generatePoints array
func generatePoints(s string) ([]Point, error) {

	points := []Point{}

	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	vals := strings.Split(s, ",")
	if len(vals) < 2 {
		return []Point{}, fmt.Errorf("Point [%v] was not well defined", s)
	}

	var x, y float64

	for idx, val := range vals {

		if idx%2 == 0 {
			x, _ = strconv.ParseFloat(val, 64)
		} else {
			y, _ = strconv.ParseFloat(val, 64)
			points = append(points, Point{x, y})
		}
	}
	return points, nil
}

// calculate the distance between two points
func distance(a Point,b Point) float64 {
	var distance float64
	distance = math.Sqrt(math.Pow(b.X-a.X ,2) + math.Pow(b.Y-a.Y ,2))
	return math.Abs(distance)
}

// verify if three points al collinear or not
func areCollinear(a,b,c Point) bool {
	var u, v Line
	var m1, m2, b1, b2 float64
	u.A = b.X - a.X
	u.B = b.Y - a.Y
	u.C = (b.X * a.Y) - (a.X * b.Y)
	v.A = c.X - b.X
	v.B = c.Y - b.Y
	v.C = (c.X * b.Y) - (b.X * c.Y)
	m1 = -u.A/u.B
	b1 = -u.C/u.B
	m2 = -v.A/v.B
	b2 = -v.C/v.B
	if m1 != m2 || b1 != b2 {
		return false
	}else{
		return true
	}
}

// getArea gets the area inside from a given shape
func getArea(points []Point) float64 {
	var area float64
	var allAreCollinear bool
	allAreCollinear = true
	area = -1
	if len(points) > 2 { // verify it is a polygon, not a line or something else
		 area = 0
		 points = append(points, points[0]) // add the first point again at the end
		for index, point := range points[:len(points)-1] {
			if index < len(points)-2 {
				if !areCollinear(point, points[index+1], points[index+2]) { // verify current point and the next two aren't collinear
					allAreCollinear = false // only with one none collinear point it is a polygon
				}
			}
			area += (point.X * points[index+1].Y) - (point.Y*points[index+1].X)
		}
		area = math.Abs(area/2)
	}
	if allAreCollinear { // verify it is a polygon
		area = -1
	}
	return area
	
}

// getPerimeter gets the perimeter from a given array of connected points
func getPerimeter(points []Point) float64 {
	var perimeter float64
	var allAreCollinear bool
	allAreCollinear = true
	perimeter = -1
	if len(points) > 2 { // verify it is a polygon, not a line or something else
		perimeter = 0
		points = append(points, points[0]) // add the first point again at the end
		for index, point := range points[:len(points)-1] {
			if index < len(points)-2 { 
				if !areCollinear(point, points[index+1], points[index+2]) { // verify current point and the next two aren't collinear
					allAreCollinear = false // only with one none collinear point it is a polygon
				}
			}
			perimeter += distance(point, points[index+1])
		}
	}
	if allAreCollinear { // verify it is a polygon
		perimeter = -1
	}
	return perimeter
}

// handler handles the web request and reponds it
func handler(w http.ResponseWriter, r *http.Request) {

	var vertices []Point
	for k, v := range r.URL.Query() {
		if k == "vertices" {
			points, err := generatePoints(v[0])
			if err != nil {
				fmt.Fprintf(w, fmt.Sprintf("error: %v", err))
				return
			}
			vertices = points
			break
		}
	}

	// Results gathering
	area := getArea(vertices)
	perimeter := getPerimeter(vertices)

	// Logging in the server side
	log.Printf("Received vertices array: %v", vertices)

	// Response construction
	response := fmt.Sprintf("Welcome Friend to the Remote Shapes Analyzer\n")
	response += fmt.Sprintf(" - Your figure has : [%v] vertices\n", len(vertices))
	response += fmt.Sprintf(" - Vertices        : %v\n", vertices)
	response += fmt.Sprintf(" - Perimeter       : %v\n", perimeter)
	response += fmt.Sprintf(" - Area            : %v\n", area)

	// Send response to client
	fmt.Fprintf(w, response)
}
