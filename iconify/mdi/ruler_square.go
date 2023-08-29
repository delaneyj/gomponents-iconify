package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5v16h6v-1.5H7V18h2v-1.5H5V15h4v-1.5H7V12h2v-1.5H5V9h4V5h1.5v4H12V7h1.5v2H15V5h1.5v4H18V7h1.5v2H21V3H5a2 2 0 0 0-2 2m3 2a1 1 0 0 1-1-1a1 1 0 0 1 1-1a1 1 0 0 1 1 1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}