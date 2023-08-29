package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GhostBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 19.206V12c0-5.523-4.477-10-10-10S2 6.477 2 12v7.206a1.727 1.727 0 0 0 2.5 1.544a2.891 2.891 0 0 1 2.896.18a2.892 2.892 0 0 0 3.208 0l.353-.234a1.881 1.881 0 0 1 2.086 0l.353.235a2.892 2.892 0 0 0 3.208 0a2.891 2.891 0 0 1 2.897-.18A1.727 1.727 0 0 0 22 19.206ZM16 10.5c0 .828-.448 1.5-1 1.5s-1-.672-1-1.5s.448-1.5 1-1.5s1 .672 1 1.5ZM9 12c.552 0 1-.672 1-1.5S9.552 9 9 9s-1 .672-1 1.5s.448 1.5 1 1.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}