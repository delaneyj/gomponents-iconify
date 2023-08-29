package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M16.5 10a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0Z"/>`),
		g.Group(children),
	)
}