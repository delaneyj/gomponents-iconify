package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewLabelOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19v-2h3l3.55-5L15 7H5v3H3V5h13.05L21 12l-4.95 7H12Zm-.225-7ZM5 20v-3H2v-2h3v-3h2v3h3v2H7v3H5Z"/>`),
		g.Group(children),
	)
}