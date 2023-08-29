package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xilinx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 18l5.241 6H5.586L.345 18l5.241-6L.345 6l5.241-6h7.655L8 6l5.241 6L8 18zM23.655 0H13.241l5.241 6l5.173-6zM13.241 24h10.414l-5.172-6l-5.242 6z"/>`),
		g.Group(children),
	)
}