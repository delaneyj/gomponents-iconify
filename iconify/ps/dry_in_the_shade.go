package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DryInTheShade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M3 43v426q0 18 12.5 30.5T45 512h342q17 0 29.5-12.5T429 469V43q0-18-12.5-30.5T387 0H45Q28 0 15.5 12.5T3 43zm98 0L45 98V43h56zM45 158L161 43h46L45 205v-47zm342 311H45V265L267 43h120v426z"/>`),
		g.Group(children),
	)
}