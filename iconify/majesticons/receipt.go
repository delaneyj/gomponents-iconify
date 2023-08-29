package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Receipt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M7 2a3 3 0 0 0-3 3v16a1 1 0 0 0 1.496.868L8.5 20.152l2.012 1.15a3 3 0 0 0 2.976 0l2.012-1.15l3.004 1.716A1 1 0 0 0 20 21V5a3 3 0 0 0-3-3H7z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}