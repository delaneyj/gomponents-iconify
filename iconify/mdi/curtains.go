package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Curtains(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 3H1V1h22v2M2 22h4c0-3-2-5-2-5c6-4 7-13 7-13H2v18M22 4h-9s1 9 7 13c0 0-2 2-2 5h4V4Z"/>`),
		g.Group(children),
	)
}