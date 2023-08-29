package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VisualRecognition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v10h-2V4H4v9.586l5-5l4.914 4.914l-1.414 1.414l-3.5-3.5l-5 5V20h13v2H2V2Zm13.547 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0ZM16 14h8v2h-3v7h-2v-7h-3v-2Z"/>`),
		g.Group(children),
	)
}