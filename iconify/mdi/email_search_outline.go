package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailSearchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 11c2.5 0 4.5 2 4.5 4.5c0 .88-.25 1.71-.69 2.4l3.08 3.1L22 22.39l-3.12-3.07c-.69.43-1.51.68-2.38.68c-2.5 0-4.5-2-4.5-4.5s2-4.5 4.5-4.5m0 2a2.5 2.5 0 0 0 0 5a2.5 2.5 0 0 0 0-5m-6 5H3V8l7.62 4.76A6.486 6.486 0 0 1 16.5 9c.27 0 .54 0 .81.06L19 8v1.5c.75.31 1.42.77 2 1.32V6c0-1.1-.9-2-2-2H3c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h8.82c-.55-.58-1-1.25-1.32-2M19 6l-8 5l-8-5h16Z"/>`),
		g.Group(children),
	)
}