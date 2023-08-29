package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GhostBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M22 19.206V12c0-5.523-4.477-10-10-10S2 6.477 2 12v7.206a1.727 1.727 0 0 0 2.5 1.544a2.891 2.891 0 0 1 2.896.18a2.892 2.892 0 0 0 3.208 0l.353-.234a1.881 1.881 0 0 1 2.086 0l.353.235a2.892 2.892 0 0 0 3.208 0a2.891 2.891 0 0 1 2.897-.18A1.727 1.727 0 0 0 22 19.206Z" opacity=".5"/><path d="M15 12c.552 0 1-.672 1-1.5S15.552 9 15 9s-1 .672-1 1.5s.448 1.5 1 1.5Zm-5-1.5c0 .828-.448 1.5-1 1.5s-1-.672-1-1.5S8.448 9 9 9s1 .672 1 1.5Z"/></g>`),
		g.Group(children),
	)
}