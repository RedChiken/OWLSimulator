package owlstruct

type TeamInfo struct {
	name  *TeamName
	match []MatchInfo
}

//total number of win match
func (team *TeamInfo) totalWin(matchInfo []MatchInfo) uint32 {
	var total uint32
	for _, value := range matchInfo {
		if value.winMatch(team.name) {
			total++
		}
	}
	return total
}

//total number of lose match
func (team *TeamInfo) totalLose(matchInfo []MatchInfo) uint32 {
	return 7 - team.totalWin(matchInfo)
}

//total number of win set
func (team *TeamInfo) totalWinSet(matchInfo []MatchInfo) uint32 {
	var total uint32
	for _, value := range matchInfo {
		total += value.win(team.name)
	}
	return total
}

//total number of lose set
func (team *TeamInfo) totalLoseSet(matchInfo []MatchInfo) uint32 {
	var total uint32
	for _, value := range matchInfo {
		total += value.lose(team.name)
	}
	return total
}

//total number of draw set
func (team *TeamInfo) totalDrawSet(matchInfo []MatchInfo) uint32 {
	var total uint32
	for _, value := range matchInfo {
		total += value.draw(team.name)
	}
	return total
}

//rate of win set
func (team *TeamInfo) winRate(matchInfo []MatchInfo) float32 {
	return float32(team.totalWinSet(matchInfo)) / float32(team.totalWinSet(matchInfo)+team.totalLoseSet(matchInfo)+team.totalDrawSet(matchInfo))
}
