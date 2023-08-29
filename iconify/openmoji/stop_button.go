package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M49.261 18A4.744 4.744 0 0 1 54 22.739v26.523A4.744 4.744 0 0 1 49.261 54H22.74A4.744 4.744 0 0 1 18 49.261V22.74A4.744 4.744 0 0 1 22.739 18H49.26m0-2H22.74A6.739 6.739 0 0 0 16 22.739v26.523A6.739 6.739 0 0 0 22.739 56h26.523A6.739 6.739 0 0 0 56 49.261V22.74A6.739 6.739 0 0 0 49.261 16z"/>`),
		g.Group(children),
	)
}