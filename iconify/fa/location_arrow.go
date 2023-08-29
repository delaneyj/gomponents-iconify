package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1401 93L761 1373q-17 35-57 35q-5 0-15-2q-22-5-35.5-22.5T640 1344V768H64q-22 0-39.5-13.5T2 719t4-42t29-30L1315 7q13-7 29-7q27 0 45 19q15 14 18.5 34.5T1401 93z"/>`),
		g.Group(children),
	)
}