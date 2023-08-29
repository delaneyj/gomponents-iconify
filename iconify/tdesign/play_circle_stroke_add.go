package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleStrokeAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0 0 18h1v2h-1C5.925 23 1 18.075 1 12S5.925 1 12 1s11 4.925 11 11v1h-2v-1a9 9 0 0 0-9-9ZM9.5 7.131L16.803 12L9.5 16.869V7.13Zm2 3.737v2.264L13.197 12L11.5 10.868ZM20 15v3h3v2h-3v3h-2v-3h-3v-2h3v-3h2Z"/>`),
		g.Group(children),
	)
}