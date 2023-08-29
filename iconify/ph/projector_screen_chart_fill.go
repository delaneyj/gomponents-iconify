package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorScreenChartFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 64V48a16 16 0 0 0-16-16H40a16 16 0 0 0-16 16v16a16 16 0 0 0 16 16v96h-8a8 8 0 0 0 0 16h88v17.38a24 24 0 1 0 16 0V192h88a8 8 0 0 0 0-16h-8V80a16 16 0 0 0 16-16Zm-128 80a8 8 0 0 1-16 0v-16a8 8 0 0 1 16 0Zm24 96a8 8 0 1 1 8-8a8 8 0 0 1-8 8Zm8-96a8 8 0 0 1-16 0v-24a8 8 0 0 1 16 0Zm32 0a8 8 0 0 1-16 0v-32a8 8 0 0 1 16 0ZM40 64V48h176v16H40Z"/>`),
		g.Group(children),
	)
}