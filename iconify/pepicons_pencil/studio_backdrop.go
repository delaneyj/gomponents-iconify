package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioBackdrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M3.5 3a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1V3Zm12 0h-11v7h11V3Z"/><path d="M2.25 2.5a.5.5 0 0 1 .5-.5h14.5a.5.5 0 0 1 0 1H2.75a.5.5 0 0 1-.5-.5Zm2.24 7.598L3.11 17h13.78l-1.38-6.902l.98-.196l1.38 6.902A1 1 0 0 1 16.89 18H3.11a1 1 0 0 1-.98-1.196l1.38-6.902l.98.196Z"/></g>`),
		g.Group(children),
	)
}