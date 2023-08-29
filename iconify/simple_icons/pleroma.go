package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pleroma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.36 0a1.868 1.868 0 0 0-1.87 1.868V24h5.964V0zm7.113 0v12h4.168a1.868 1.868 0 0 0 1.868-1.868V0zm0 18.036V24h4.168a1.868 1.868 0 0 0 1.868-1.868v-4.096Z"/>`),
		g.Group(children),
	)
}