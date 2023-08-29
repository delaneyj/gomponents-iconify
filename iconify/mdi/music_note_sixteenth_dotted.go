package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteSixteenthDotted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 18.5a1.5 1.5 0 1 1-1.5-1.5a1.5 1.5 0 0 1 1.5 1.5M18 7V3h-6v10.55A4 4 0 1 0 14 17v-6h4V8h-4V7Z"/>`),
		g.Group(children),
	)
}