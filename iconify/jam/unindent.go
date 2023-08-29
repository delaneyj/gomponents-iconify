package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unindent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 0h12a1 1 0 0 1 0 2H1a1 1 0 1 1 0-2zm0 8h12a1 1 0 0 1 0 2H1a1 1 0 1 1 0-2zm6-4h6a1 1 0 0 1 0 2H7a1 1 0 1 1 0-2zm-6.44.143l2.057-1.234a1 1 0 0 1 1.515.857v2.468a1 1 0 0 1-1.515.857L.561 5.857a1 1 0 0 1 0-1.714z"/>`),
		g.Group(children),
	)
}