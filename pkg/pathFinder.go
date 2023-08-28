package pkg

import (
	"lemin/models"
	"time"
)

func FindPaths(rooms map[string]models.Room, start string, end string, path []string, paths *[][]string, timeout time.Duration) [][]string {
	// Début du délai de temps
	startTime := time.Now()

	// Marquer le nœud actuel comme visité
	buffer := rooms[start]
	buffer.Visited = true
	rooms[start] = buffer

	// Ajouter le nœud actuel au chemin en cours de construction
	path = append(path, start)

	// Si le nœud actuel est le nœud de fin, ajouter le chemin complet à paths
	if start == end {
		*paths = append(*paths, append([]string{}, path...))
	} else {
		// Parcourir les voisins du nœud actuel
		for _, neighbor := range rooms[start].Links {
			// Vérifier si le délai de temps est dépassé
			if time.Since(startTime) >= timeout {
				return *paths
			}

			// Vérifier si le voisin n'a pas été visité
			if !rooms[neighbor].Visited {
				// Récursivement, chercher les chemins à partir du voisin vers le nœud de fin
				FindPaths(rooms, neighbor, end, append([]string{}, path...), paths, timeout)
			}
		}
	}

	// Remettre à zéro la marque du nœud actuel pour les futures recherches
	buffer.Visited = false
	rooms[start] = buffer

	return *paths
}

func IsOptimal(path1, path2 []string) bool {
	for i, v := range path1 {
		if i == 0 || i == len(path1)-1 {
			continue
		}
		for j, r := range path2 {
			if j == 0 || j == len(path2)-1 {
				continue
			}
			if v == r {
				return false
			}
		}
	}
	return true
}
func Minmax(paths [][]string) [][]string {
	for i := 0; i < len(paths); i++ {
		for j, k := range paths {
			if i != j && len(paths[i]) <= len(k) {
				tmp := paths[i]
				paths[i] = paths[j]
				paths[j] = tmp
			}
		}
	}
	return paths
}
func BestPath(PathGroups [][][]string, ants int) [][]models.Room {
	var min int
	var bestGroup [][]string
	for _, group := range PathGroups {
		maxPathLength := len(group[len(group)-1])
		neededFill := 0
		// Calculer la quantité totale de remplissage nécessaire
		for _, path := range group {
			neededFill += maxPathLength - len(path)
		}
		leftAnts := ants - neededFill
		// Si toutes les fourmis sont utilisées, passer au prochain groupe
		if leftAnts == 0 {
			continue
		}
		// Calculer la hauteur en répartissant les fourmis restantes entre les chemins
		height := maxPathLength + leftAnts/len(group)
		if leftAnts%len(group) != 0 {
			height++
		}
		// Mettre à jour la meilleure hauteur et le meilleur groupe si nécessaire
		if min == 0 || height < min {
			min = height
			bestGroup = group
		}
	}
	var pathMap [][]models.Room
	for _, v := range bestGroup {
		var path []models.Room
		for _, r := range v {
			room := models.Room{
				Path: r,
			}
			path = append(path, room)
		}
		pathMap = append(pathMap, path)
	}
	return pathMap
}
func Filter(paths [][]string) [][][]string {
	// parcourrir l'ensemble des paths et trier en fonction de la taille du plus petit paths au plus grand de maniere croissant
	paths = Minmax(paths)
	//enlever toutes les collisions dans les paths
	var prePath = AllPathCollision(paths)
	return prePath
}
func AllPathCollision(paths [][]string) [][][]string {
	var filteredPaths [][]string
	var allCollision [][][]string
	for _, d := range paths {
		filteredPaths = append(filteredPaths, d)
		var actif = false
		var tab []string
		for i, k := range paths {
			if i == 0 {
				continue
			}
			for _, v := range filteredPaths {
				if IsOptimal(v, k) {
					actif = true
					tab = k
				} else {
					actif = false
					break
				}
			}
			if actif {
				filteredPaths = append(filteredPaths, tab)
				actif = false
				tab = []string{}
			}
		}
		allCollision = append(allCollision, filteredPaths)
		filteredPaths = [][]string{}
	}
	return allCollision
}

func FilterArray(arr *[]string, target string) {
	var newArray []string
	for _, item := range *arr {
		if item == target {
			newArray = append(newArray, item)
			*arr = newArray
		}
	}
}

func RemoveUnnecessaryLinks(rooms map[string]models.Room, end string) {
	for i := range rooms {
		buffer := rooms[i]
		FilterArray(&buffer.Links, end)
		rooms[i] = buffer
	}
}
