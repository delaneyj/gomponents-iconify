package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.828 1.416l5.755 5.755L7.755 22H2v-5.756L16.828 1.416Zm0 8.681l2.927-2.926l-2.927-2.927l-2.926 2.927l2.926 2.926Zm-4.34-1.512L4 17.074V20h2.926l8.488-8.488l-2.926-2.927Z"/>`),
		g.Group(children),
	)
}