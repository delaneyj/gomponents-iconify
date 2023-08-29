package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReligionShintoReligionGateCultureShintoJapanJapaneseShrine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 4a.5.5 0 0 1-.5.5H1A.5.5 0 0 1 .5 4V.5a11.55 11.55 0 0 0 13 0ZM3 4.5v9m8-9v9M.5 8h13M7 4.5V8"/>`),
		g.Group(children),
	)
}