package pkg

import (
	"lemin/models"
	"strings"
)

func EdgeLinks(room string, tunnels map[string]models.Tunnel) []string {
	var Links []string
	for key := range tunnels {
		if strings.Contains(key, room) {
			Links = append(Links, RemoveSubstrings(key, []string{room, "-"}))
		}
	}
	return Links
}

func RemoveSubstrings(input string, substrings []string) string {
	result := input
	for _, substr := range substrings {
		result = strings.Replace(result, substr, "", -1)
	}
	return result
}

func FindLinks(tunnels map[string]models.Tunnel, rooms map[string]models.Room) {
	for key := range tunnels {
		links := strings.Split(key, "-")
		sourceRoom := rooms[links[0]]
		targetRoom := rooms[links[1]]
		sourceRoom.UpdateLinks(links[1])
		targetRoom.UpdateLinks(links[0])
		
		rooms[links[0]] = sourceRoom
		rooms[links[1]] = targetRoom
	}
}

func BuildFlowNetwork(rooms map[string]models.Room, tunnels map[string]models.Tunnel) map[string]map[string]int {
	flowNetwork := make(map[string]map[string]int)

	for roomName := range rooms {
		flowNetwork[roomName] = make(map[string]int)
	}

	for _, tunnel := range tunnels {
		// Pour chaque tunnel, ajouter la capacité résiduelle de 1 (ou ce que tu veux) dans les deux directions
		flowNetwork[tunnel.From][tunnel.To] = 1
		flowNetwork[tunnel.To][tunnel.From] = 1
	}

	return flowNetwork
}
