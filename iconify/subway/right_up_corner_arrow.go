package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightUpCornerArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m186.1.1l-93 93.1H349L0 442.2L69.8 512l349-349v255.9l93.1-93L512 0z"/>`),
		g.Group(children),
	)
}