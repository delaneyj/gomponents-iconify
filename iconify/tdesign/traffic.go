package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Traffic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 2h12v4h3v2h-3v3h3v2h-3v3h3v2h-3v4H6v-4H3v-2h3v-3H3v-2h3V8H3V6h3V2Zm2 2v16h8V4H8Zm2 3a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm0 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm0 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z"/>`),
		g.Group(children),
	)
}