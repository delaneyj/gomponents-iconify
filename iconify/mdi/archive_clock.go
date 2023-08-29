package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 6H2V2h18v4m-3.5 6H15v5l3.61 2.16l.75-1.22l-2.86-1.69V12m6.5 4c0 3.87-3.13 7-7 7c-2.38 0-4.47-1.19-5.74-3H3V7h16v2.68c2.36 1.13 4 3.53 4 6.32M8 12h2.26c.57-.81 1.3-1.5 2.15-2H8.5c-.28 0-.5.22-.5.5V12m13 4c0-2.76-2.24-5-5-5s-5 2.24-5 5s2.24 5 5 5s5-2.24 5-5Z"/>`),
		g.Group(children),
	)
}