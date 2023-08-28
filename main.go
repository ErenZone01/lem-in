package main

import (
	"fmt"
	"lemin/models"
	"lemin/pkg"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run . file.txt")
		return
	}
	rooms := make(map[string]models.Room)
	tunnels := make(map[string]models.Tunnel)
	ants, start, end, input := pkg.DataParser(args[1], rooms, tunnels)
	numAnts, _ := strconv.Atoi(ants)
	startLinks := pkg.EdgeLinks(start, tunnels)
	endLinks := pkg.EdgeLinks(end, tunnels)
	pkg.DataValidation(start, end, startLinks, endLinks, rooms, tunnels)
	pkg.FindLinks(tunnels, rooms) //Pour BFS cette ligne n'est pas n'est pas requise
	fmt.Println(ants)
	for i := 0; i < len(input); i++ {
		fmt.Printf("%v\n", input[i])
	}
	fmt.Println()
	allPaths := pkg.FindPaths(rooms, start, end, []string{}, &[][]string{}, 150000000)
	bestPaths := pkg.BestPath(pkg.Filter(allPaths), numAnts)
	pkg.Move(bestPaths, numAnts)
}
