package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpScreenshot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 1v22h14V1H5zm12 17H7V6h10v12zM9.5 8.5H12V7H8v4h1.5V8.5zM12 17h4v-4h-1.5v2.5H12V17z"/>`),
		g.Group(children),
	)
}