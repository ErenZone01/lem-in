package pkg

import (
	"bufio"
	"fmt"
	"lemin/models"
	"os"
	"strconv"
	"strings"
)

func IsValidRoom(line string) bool {
	// Diviser la ligne en éléments séparés par un espace
	elements := strings.Split(line, " ")
	if len(elements) != 3 {
		return false
	}

	firstElement := elements[0]
	secondElement := elements[1]
	thirdElement := elements[2]

	// Vérifier les critères
	if strings.HasPrefix(firstElement, "L") || strings.HasPrefix(firstElement, "#") {
		return false
	}

	if (!IsPositiveInteger(secondElement) || !IsPositiveInteger(thirdElement)) && (!IsFloat(secondElement) || !IsFloat(thirdElement)) {
		return false
	}

	return true
}

func IsComment(element string) bool {
	return strings.HasPrefix(element, "#") && !strings.HasPrefix(element, "##")
}

func IsStart(element string) bool {
	return element == "##start"
}

func IsEnd(element string) bool {
	return element == "##end"
}

func IsPositiveInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsValidTunnel(tunnel string) bool {
	parts := strings.Split(tunnel, "-")
	if len(parts) != 2 {
		return false
	}
	if len(parts) == 2 {
		if strings.Contains(parts[0], " ") || strings.Contains(parts[1], " ") {
			return false
		}
	}
	return true
}

func HasDuplicateRoom(rooms map[string]models.Room) bool {
	seen := make(map[string]bool)
	for _, room := range rooms {
		coords := room.X + " " + room.Y
		if seen[coords] {
			return true // Les coordonnées sont dupliquées
		}
		seen[coords] = true // Marquer les coordonnées comme vues
	}

	return false // Aucune coordonnée dupliquée trouvée
}

func DataParser(path string, rooms map[string]models.Room, tunnels map[string]models.Tunnel) (numants, startPoint, endPoint string, input []string) {
	var lines []string
	var room models.Room
	var tunnel models.Tunnel
	start := false
	end := false
	numStart, numEnd := 0, 0
	file, err := os.Open(path)
	var foundStart bool
	var foundEnd bool
	source := ""
	target := ""
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		os.Exit(0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ants := ""
lco:
	if scanner.Scan() {
		if IsComment(scanner.Text()) {
			goto lco
		}
		ants = scanner.Text()
	}
	antsNum, _ := strconv.Atoi(ants)
	if !IsPositiveInteger(ants) || antsNum == 0 || strings.HasPrefix(ants, "-") {
		fmt.Println("ERROR: invalid data format, invalid number of Ants")
		os.Exit(0)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if !IsComment(line) && !IsStart(line) && !IsValidRoom(line) && !IsEnd(line) && !IsValidTunnel(line) {
			fmt.Println("ERROR: invalid data format")
			os.Exit(0)
		}
		if IsValidRoom(line) {
			if foundStart {
				source = strings.Split(line, " ")[0]
				foundStart = false
			}
			if foundEnd {
				target = strings.Split(line, " ")[0]
				foundEnd = false
			}
			room.InitRoom(strings.Split(line, " "))
			_, exist := rooms[strings.Split(line, " ")[0]]
			if exist {
				fmt.Println("ERROR: invalid data format, duplicated room founded")
				os.Exit(0)
			}

			rooms[strings.Split(line, " ")[0]] = room
		}
		if IsValidTunnel(line) {
			tunnel.InitTunnel(strings.Split(line, "-"))
			_, exist := tunnels[line]
			if tunnel.From == tunnel.To || exist {
				fmt.Println("ERROR: invalid data format, duplicated or Auto-linked tunnel")
				os.Exit(0)
			}
			tunnels[line] = tunnel
		}
		if IsStart(line) {
			start = true
			numStart++
			foundStart = true
		}
		if IsEnd(line) {
			end = true
			numEnd++
			foundEnd = true
		}
		lines = append(lines, line)
	}

	// Vérifier les erreurs lors de la lecture
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
	}
	if !start || numStart != 1 {
		fmt.Println("ERROR: invalid data format, no or to much start point found in the file")
		os.Exit(0)
	}
	if !end || numEnd != 1 {
		fmt.Println("ERROR: invalid data format, no or to much end point found in the file")
		os.Exit(0)
	}
	if HasDuplicateRoom(rooms) {
		fmt.Println("ERROR: invalid data format, rooms with same coordinated")
		os.Exit(0)
	}
	return ants, source, target, lines
}

func FindUnknownRoom(tunnels map[string]models.Tunnel, rooms map[string]models.Room) bool {
	for _, tunel := range tunnels {
		if rooms[tunel.From].X == "" || rooms[tunel.To].Y == "" {
			return true
		}
	}
	return false
}

func DataValidation(start, end string, startLinks, endLinks []string, rooms map[string]models.Room, tunnels map[string]models.Tunnel) {
	if start == "" || end == "" {
		fmt.Println("ERROR: invalid data format")
		os.Exit(0)
	}
	if len(startLinks) == 0 || len(endLinks) == 0 {
		fmt.Println("ERROR: invalid data format, no room linked with edges rooms")
		os.Exit(0)
	}
	if FindUnknownRoom(tunnels, rooms) {
		fmt.Println("ERROR: invalid data format, unknown room in the graph")
		os.Exit(0)
	}
}
