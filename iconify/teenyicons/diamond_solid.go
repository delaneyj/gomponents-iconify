package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M11.783 0H3.217L.07 5.243a.5.5 0 0 0 .034.564l7 9a.5.5 0 0 0 .79 0l7-9a.5.5 0 0 0 .034-.564L11.783 0Z"/>`),
		g.Group(children),
	)
}