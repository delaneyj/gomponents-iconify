package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 2H3v3H0v2h3v3h2V7h3V5H5V2zm12 1h-7v2h5v2h5v12H5v-7H3v9h19V5h-5V3zm-7 6h4v2h2v4h-2v2h-4v-2h4v-4h-4V9zm-2 2h2v4H8v-4z"/>`),
		g.Group(children),
	)
}