package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoLabelOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 12l-4.95 7H3V5h13.05L21 12Zm-2.45 0L15 7H5v10h10l3.55-5ZM5 12v5V7v5Zm5.525 4l1.25-2.75l2.75-1.25l-2.75-1.25L10.525 8l-1.25 2.75L6.525 12l2.75 1.25l1.25 2.75Z"/>`),
		g.Group(children),
	)
}