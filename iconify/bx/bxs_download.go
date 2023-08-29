package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M19 9h-4V3H9v6H5l7 8zM4 19h16v2H4z" fill="currentColor"/>`),
		g.Group(children),
	)
}