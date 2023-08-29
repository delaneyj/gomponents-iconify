package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.25 19a.75.75 0 0 1 0 1.5h-6.5a.75.75 0 0 1 0-1.5h6.5Zm4-6a.75.75 0 0 1 0 1.5H6.75a.75.75 0 0 1 0-1.5h14.5Zm3-6a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1 0-1.5h20.5Z"/>`),
		g.Group(children),
	)
}