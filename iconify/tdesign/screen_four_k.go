package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenFourK(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1V3Zm2 2v14h18V5H3Zm12 3v2.048l1-.708V8h2v1.598a1.5 1.5 0 0 1-.633 1.225L15.703 12l1.664 1.177A1.5 1.5 0 0 1 18 14.402V16h-2v-1.34l-1-.708V16h-2V8h2ZM7 8v3.429h2V8h2v8H9v-2.571H7a2 2 0 0 1-2-2V8h2Z"/>`),
		g.Group(children),
	)
}