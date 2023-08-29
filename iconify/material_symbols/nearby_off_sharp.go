package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NearbyOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.8 16L17 14.2l2.2-2.2L12 4.8L9.8 7L8 5.2l4-4L22.8 12l-4 4Zm-6.825 6.775l-10.75-10.75l3.975-4L1.4 4.2l1.4-1.4l18.4 18.4l-1.375 1.375L16.05 18.8l-4.075 3.975ZM12 19.2l2.25-2.2l-1.425-1.425L12 16.4l-4.375-4.375l.825-.825l-1.4-1.4l-2.225 2.225L12 19.2Zm3.6-6.4l-4.4-4.4l.8-.8l4.4 4.4l-.8.8Z"/>`),
		g.Group(children),
	)
}