package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Down(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1965 1101l-941 941l-941-941l90-90l787 787V0h128v1798l787-787l90 90z"/>`),
		g.Group(children),
	)
}