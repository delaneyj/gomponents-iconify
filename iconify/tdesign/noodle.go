package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Noodle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 3h7v2h-1v5h1v2c0 5.523-4.477 10-10 10S2 17.523 2 12v-2h1a6 6 0 1 1 12 0h1V5h-1V3Zm3 2v5h1V5h-1ZM4 12a8 8 0 1 0 16 0H4Zm1-2h1a3 3 0 0 1 6 0h1a4 4 0 0 0-8 0Zm5 0a1 1 0 0 0-2 0h2Z"/>`),
		g.Group(children),
	)
}