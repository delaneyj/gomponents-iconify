package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Triangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.934 20.006l-9.5-16.51a.521.521 0 0 0-.868 0l-9.5 16.51a.5.5 0 0 0 .434.749h19a.5.5 0 0 0 .434-.75zm-18.57-.251L12 4.748l8.636 15.007H3.364z"/>`),
		g.Group(children),
	)
}