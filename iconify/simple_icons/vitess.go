package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vitess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.206 1.045l-7.217 13.186L4.817 1.045H0l11.904 21.91L24 1.045h-4.794Z"/>`),
		g.Group(children),
	)
}