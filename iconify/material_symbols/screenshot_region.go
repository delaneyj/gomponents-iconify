package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenshotRegion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 22v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM5 19v-5h2v3h3v2H5Zm0-9V5h5v2H7v3H5Zm12 0V7h-3V5h5v5h-2Z"/>`),
		g.Group(children),
	)
}