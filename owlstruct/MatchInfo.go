package owlstruct

type MatchInfo struct {
	matchdate *Date
	homeStand *TeamName
	home      *TeamName
	away      *TeamName
	homeScore uint32
	awayScore uint32
	drawScore uint32
}

//opponent team name
func (match *MatchInfo) opponent(team *TeamName) *TeamName {
	if match.home == team {
		return match.away
	}
	return match.home
}

//weather the team win this match or not
func (match *MatchInfo) winMatch(team *TeamName) bool {
	if match.home == team {
		return match.homeScore > match.awayScore
	}
	return match.awayScore > match.homeScore
}

//the number of set which the team win
func (match *MatchInfo) win(team *TeamName) uint32 {
	var score uint32
	if match.home == team {
		score = match.homeScore
	} else if match.away == team {
		score = match.awayScore
	} else {
		score = 5
	}
	return score
}

//the number of set which the team lose
func (match *MatchInfo) lose(team *TeamName) uint32 {
	var score uint32
	if match.home == team {
		score = match.awayScore
	} else if match.away == team {
		score = match.homeScore
	} else {
		score = 5
	}
	return score
}

//the number of set which the team draw
func (match *MatchInfo) draw(team *TeamName) uint32 {
	if (match.home == team) || (match.away == team) {
		return match.drawScore
	}
	return 5
}
