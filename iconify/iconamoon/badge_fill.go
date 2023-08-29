package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BadgeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h7.172a3 3 0 0 1 2.12.879l7 7a3 3 0 0 1 0 4.242l-5.17 5.172a3 3 0 0 1-4.243 0l-7-7A3 3 0 0 1 3 11.172V4Zm6 3.5A1.5 1.5 0 0 0 7.5 9v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9a1.5 1.5 0 0 0-1.5-1.5H9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}