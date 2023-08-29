package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleBigFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M19.5 10a9.5 9.5 0 1 1-19 0a9.5 9.5 0 0 1 19 0Z"/>`),
		g.Group(children),
	)
}