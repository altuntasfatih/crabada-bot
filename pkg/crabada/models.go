package crabada

import "fmt"

type GetMineHistoryResponse struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
	Result    struct {
		Limit int         `json:"limit"`
		Data  []*MineInfo `json:"data"`
	} `json:"result"`
}
type MineInfo struct {
	GameId    int   `json:"game_id"`
	StartTime int64 `json:"start_time"`
	Round     int   `json:"round"`
}

type GetTeamsResponse struct {
	ErrorCode string       `json:"error_code"`
	Message   string       `json:"message"`
	Result    *TeamsResult `json:"result"`
}

type TeamsResult struct {
	TeamSize int     `json:"team_size"`
	Teams    []*Team `json:"data"`
	Message  string
}

type Team struct {
	CrabadaIdOne   int    `json:"crabada_id_1"`
	CrabadaIdTwo   int    `json:"crabada_id_2"`
	CrabadaIdThree int    `json:"crabada_id_3"`
	TeamId         int    `json:"team_id"`
	BattlePoint    int    `json:"battle_point"`
	MinePoint      int    `json:"mine_point"`
	Status         string `json:"status"`
	GameType       string `json:"game_type"`
	GameRound      int    `json:"game_round"`
	ProcessStatus  string `json:"process_status"`
}

func (t *Team) IsFreeForAttack() bool {
	return t.Status == "AVAILABLE" && t.CrabadaIdOne > 0 && t.CrabadaIdTwo > 0 && t.CrabadaIdThree > 0
}

func (t *Team) ToString() string {
	return fmt.Sprintf("Team{Id: %d, BattlePoint: %d, MinePoint: %d, Status: %s }", t.TeamId, t.BattlePoint, t.MinePoint, t.Status)
}
