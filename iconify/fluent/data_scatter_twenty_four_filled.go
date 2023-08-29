package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataScatterTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 3.75a.75.75 0 0 1 1.5 0V19.5h15.75a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1-.75-.75V3.75ZM14 7a3 3 0 1 1 6 0a3 3 0 0 1-6 0ZM9 6a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm6 6a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z"/>`),
		g.Group(children),
	)
}