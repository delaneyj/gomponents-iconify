package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eyeglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4H6L3 14M16 4h2l3 10m-11 2h4m7 .5a3.5 3.5 0 0 1-7 0V14h7v2.5m-11 0a3.5 3.5 0 0 1-7 0V14h7v2.5"/>`),
		g.Group(children),
	)
}