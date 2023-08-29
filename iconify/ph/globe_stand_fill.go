package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeStandFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M48 96a80 80 0 1 1 80 80a80.09 80.09 0 0 1-80-80Zm146.46 69.28A96 96 0 0 1 58.72 29.54a8 8 0 1 0-11.54-11.08A112 112 0 0 0 120 207.71V224H96a8 8 0 0 0 0 16h64a8 8 0 0 0 0-16h-24v-16.28a111.21 111.21 0 0 0 69.54-30.9a8 8 0 1 0-11.08-11.54Z"/>`),
		g.Group(children),
	)
}