package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftUpCornerArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M163 93.2h255.9L325.9.1L0 0l.1 325.9l93.1 93V163l349 349l69.8-69.8z"/>`),
		g.Group(children),
	)
}