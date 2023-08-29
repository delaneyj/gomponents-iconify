package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenshotMonitorSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 16h4v-4h-1.5v2.5H15V16ZM5 10h1.5V7.5H9V6H5v4Zm3 11v-2H2V3h20v16h-6v2H8Z"/>`),
		g.Group(children),
	)
}