package owlstruct

type TeamInfo struct {
	name    *TeamName
	match   []MatchInfo
	picture string
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
	return 28 - team.totalWin(matchInfo)
}

//total number of win set
func (team *TeamInfo) totalWinSet(matchInfo []MatchInfo) uint32 {
	var total uint32
	var winset uint32
	for _, value := range matchInfo {
		winset = value.win(team.name)
		if winset < 5 {
			total += winset
		}
	}
	return total
}

//total number of lose set
func (team *TeamInfo) totalLoseSet(matchInfo []MatchInfo) uint32 {
	var total uint32
	var loseset uint32
	for _, value := range matchInfo {
		loseset = value.lose(team.name)
		if loseset < 5 {
			total += loseset
		}
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
	return float32(team.totalWin(matchInfo)) / float32(team.totalWin(matchInfo)+team.totalLose(matchInfo))
}

func (team *TeamInfo) Compare(another *TeamInfo, matchInfo []MatchInfo) bool {
	if team.totalWin(matchInfo) == another.totalWin(matchInfo) {
		if team.winRate(matchInfo) == another.winRate(matchInfo) {
			return team.totalWinSet(matchInfo) > another.totalWinSet(matchInfo)
		}
		return team.winRate(matchInfo) == another.winRate(matchInfo)
	}
	return team.totalWin(matchInfo) > another.totalWin(matchInfo)
}
