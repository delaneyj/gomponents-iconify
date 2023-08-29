package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsMessageX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M20 2H4c-1.103 0-2 .894-2 1.992v12.016C2 17.106 2.897 18 4 18h3v4l6.351-4H20c1.103 0 2-.894 2-1.992V3.992A1.998 1.998 0 0 0 20 2zm-3.293 11.293l-1.414 1.414L12 11.414l-3.293 3.293l-1.414-1.414L10.586 10L7.293 6.707l1.414-1.414L12 8.586l3.293-3.293l1.414 1.414L13.414 10l3.293 3.293z" fill="currentColor"/>`),
		g.Group(children),
	)
}