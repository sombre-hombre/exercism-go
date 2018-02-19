package tournament

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"
)

// Tally the results of a small football competition.
func Tally(reader io.Reader, buf io.Writer) error {
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'
	csvReader.Comment = '#'
	csvReader.FieldsPerRecord = 3

	teams := make(map[string]*team)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		err = parseRecord(record, teams)
		if err != nil {
			return err
		}
	}

	printTable(getScoreTable(teams), buf)

	return nil
}

func printTable(table []team, buf io.Writer) {
	fmt.Fprint(buf, "Team                           | MP |  W |  D |  L |  P\n")
	for _, t := range table {
		fmt.Fprintf(buf, "%-30s |%3d |%3d |%3d |%3d |%3d\n",
			t.name, t.matches, t.wins, t.draws, t.losses, t.points)
	}
}

// getScoreTable returns sorted score table
func getScoreTable(teams map[string]*team) []team {
	table := make([]team, len(teams))
	var i int
	for _, t := range teams {
		table[i] = *t
		i++
	}

	sort.Slice(table, func(i, j int) bool {
		switch {
		case table[i].points > table[j].points:
			return true
		case table[i].points < table[j].points:
			return false
		default:
			return table[i].name < table[j].name
		}
	})

	return table
}

func parseRecord(rec []string, teams map[string]*team) error {
	team1, found := teams[rec[0]]
	if !found {
		team1 = &team{name: rec[0]}
		teams[team1.name] = team1
	}
	team2, found := teams[rec[1]]
	if !found {
		team2 = &team{name: rec[1]}
		teams[team2.name] = team2
	}

	switch rec[2] {
	case "win":
		team1.Win()
		team2.Loss()
		break
	case "draw":
		team1.Draw()
		team2.Draw()
		break
	case "loss":
		team1.Loss()
		team2.Win()
		break
	default:
		return errors.New("Unknown outcome")
	}

	return nil
}

type team struct {
	name                                 string
	matches, wins, draws, losses, points int
}

func (t *team) Win() {
	t.matches++
	t.wins++
	t.points += 3
}

func (t *team) Draw() {
	t.matches++
	t.draws++
	t.points++
}

func (t *team) Loss() {
	t.matches++
	t.losses++
}
