package pkg

import (
	"fmt"
	"lemin/models"
	"sort"
	"strconv"
	"strings"
)

func Move(paths [][]models.Room, ant int) {
	round := []string{}
	tab := Dispatching(paths, ant)
	for i, v := range tab {
		for j := len(v) - 1; j >= 0; j-- {
			tab[i] = append(tab[i], v[j])
			tab[i] = tab[i][1:]
		}
	}
	Release(paths, tab)
	for _, v := range paths {
		for _, r := range v {
			if r.Ant != "" {
				round = append(round, r.Ant+"-"+r.Path)
			}
		}
	}
	SortListByNumber(round)
	fmt.Println(strings.Join(round, " "))
	round = nil
	for ant > 0 {
		for i, path := range paths {
			for j := len(path) - 1; j > 0; j-- {
				if j != len(path)-1 && path[j].Ant != "" {
					paths[i][j+1].Ant = paths[i][j].Ant
					paths[i][j].Ant = ""
				}
				if j == len(path)-1 && path[j].Ant != "" {
					paths[i][j].Ant = ""
					ant--
				}
			}
		}
		Release(paths, tab)
		for _, v := range paths {
			for _, r := range v {
				if r.Ant != "" {
					round = append(round, r.Ant+"-"+r.Path)
				}
			}
		}
		if ant != 0 {
			SortListByNumber(round)
			fmt.Println(strings.Join(round, " "))
			round = nil
		}
	}
}
func Release(paths [][]models.Room, tab map[int][]string) {
	for i, path := range paths {
		if path[1].Ant == "" && len(tab[i]) > 0 {
			paths[i][1].Ant = tab[i][0]
			tab[i] = tab[i][1:]
		}
	}
}
func Dispatching(paths [][]models.Room, ant int) map[int][]string {
	tab := make(map[int][]string)
	for ant > 0 {
		for i := 0; i < len(paths); i++ {
			if i != len(paths)-1 && len(paths[i])+len(tab[i]) > len(paths[i+1])+len(tab[i+1]) {
				continue
			}
			if i == len(paths)-1 && len(paths[i])+len(tab[i]) > len(paths[0])+len(tab[0]) {
				i = 0
			}
			if ant > 0 {
				teste := "L" + strconv.Itoa(ant)
				tab[i] = append(tab[i], teste)
				ant--
			}
		}
	}
	return tab
}

func SortListByNumber(list []string) {
	sort.Slice(list, func(i, j int) bool {
		num1 := ExtractNumber(list[i])
		num2 := ExtractNumber(list[j])
		return num1 < num2
	})
}

func ExtractNumber(s string) int {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return 0
	}
	numStr := parts[0][1:] // Remove the "L" prefix
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0
	}
	return num
}

func CheckStringInArray(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}
