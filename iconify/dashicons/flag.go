package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 18V3H3v15h2zm1-6V4c3-1 7 1 11 0v8c-3 1.27-8-1-11 0z"/>`),
		g.Group(children),
	)
}