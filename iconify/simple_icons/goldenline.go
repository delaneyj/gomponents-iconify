package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Goldenline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.997 24a11.995 11.995 0 0 0 11.949-13.04h-6.781v2.943h1.226a6.667 6.667 0 1 1-.114-4.156h5.509A11.995 11.995 0 1 0 12 23.991z"/>`),
		g.Group(children),
	)
}