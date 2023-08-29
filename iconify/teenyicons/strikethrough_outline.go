package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrikethroughOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.5 10A3.5 3.5 0 0 0 7 13.5h1.487a3.013 3.013 0 0 0 3.013-3.013c0-1.205-.892-2.512-2-2.987M6.08 6.177A2.607 2.607 0 0 1 4.5 3.781A2.281 2.281 0 0 1 6.781 1.5H8A2.5 2.5 0 0 1 10.5 4M2 7.5h11"/>`),
		g.Group(children),
	)
}