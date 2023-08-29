package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenshotFrameSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 7V2h5v2H7v3H5Zm0 15v-5h2v3h3v2H5ZM17 7V4h-3V2h5v5h-2Zm-3 15v-2h3v-3h2v5h-5Z"/>`),
		g.Group(children),
	)
}