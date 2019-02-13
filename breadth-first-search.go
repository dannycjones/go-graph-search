package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type color int

const (
	white color = 0 // Unseen
	gray  color = 1 // Active
	black color = 2 // Seen
)

type vertex struct {
	index         int
	color         color
	predecessor   *vertex
	distance      int
	adjacencyList []*vertex
}

// 'toString' function for vertex
func (v vertex) String() string {
	s := fmt.Sprintf("vertex %d { color = %d ", v.index, v.color)
	if v.predecessor != nil {
		s += fmt.Sprintf("predecessor = %d ", v.predecessor.index)
	}
	s += fmt.Sprintf("distance = %d }\n", v.distance)
	return s
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Program requires first argument - file path for adjacency lists")
		return
	}
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error occurred opening file", err)
	}

	scanner := bufio.NewScanner(file)

	var vertices []vertex

	if scanner.Scan() {
		if n, err := strconv.Atoi(scanner.Text()); err == nil && n >= 0 {
			vertices = make([]vertex, n)
		} else {
			panic("First line did not contain valid number of vertices")
		}
	}

	for i := 0; i < len(vertices); i++ {
		vertices[i].index = i
		vertices[i].color = white
		vertices[i].distance = -1
	}

	for vIdx := 0; scanner.Scan(); {
		adjacencyStr := strings.Split(scanner.Text(), ",")

		for _, adjacentIndexStr := range adjacencyStr {
			// For each adjacent vertex
			if adjacentIndex, err := strconv.Atoi(adjacentIndexStr); err == nil {
				vertices[vIdx].adjacencyList = append(vertices[vIdx].adjacencyList, &vertices[adjacentIndex])
			}
		}

		vIdx++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	vertices[0].color = gray
	vertices[0].distance = 0

	queue := make([]*vertex, 0)
	queue = append(queue, &vertices[0])

	for len(queue) > 0 {
		// Go through each vertex in the queue, add any unseen to the queue
		var u *vertex
		queue, u = remove(queue)
		for _, v := range u.adjacencyList {
			if v.color == white {
				v.color = gray
				v.distance = u.distance + 1
				v.predecessor = u
				queue = append(queue, v)
			}
		}
		u.color = black // Vertex now fully explored
	}

	fmt.Println(vertices)
}

func remove(queue []*vertex) ([]*vertex, *vertex) {
	element := queue[0]
	queue = queue[1:]
	return queue, element
}
