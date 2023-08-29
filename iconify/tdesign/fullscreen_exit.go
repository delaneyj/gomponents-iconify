package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullscreenExit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.414 5l-4.5 4.5H20v2h-7.5V4h2v4.086l4.5-4.5L20.414 5ZM4 12.5h7.5V20h-2v-4.086l-4.5 4.5L3.586 19l4.5-4.5H4v-2Z"/>`),
		g.Group(children),
	)
}