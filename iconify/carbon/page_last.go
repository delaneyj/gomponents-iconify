package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageLast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18 16L8 26l-1.4-1.4l8.6-8.6l-8.6-8.6L8 6zm4-12h2v24h-2z"/>`),
		g.Group(children),
	)
}