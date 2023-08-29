package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AttachSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 12A6.75 6.75 0 0 1 9 5.25h8a4.75 4.75 0 1 1 0 9.5H8.857a2.75 2.75 0 1 1 0-5.5h7.857a.75.75 0 0 1 0 1.5H8.857a1.25 1.25 0 1 0 0 2.5H17a3.25 3.25 0 0 0 0-6.5H9a5.25 5.25 0 1 0 0 10.5h7.714a.75.75 0 0 1 0 1.5H9A6.75 6.75 0 0 1 2.25 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}