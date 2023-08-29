package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShowerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M64 236a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm20-44a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm-64 0a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm32-32a12 12 0 1 0 12 12a12 12 0 0 0-12-12ZM248 32h-28.69A15.86 15.86 0 0 0 208 36.69l-27.86 27.85L53.38 86.19a16 16 0 0 0-8.69 27.1l98 98a16 16 0 0 0 27.09-8.66l21.68-126.77L219.31 48H248a8 8 0 0 0 0-16Z"/>`),
		g.Group(children),
	)
}