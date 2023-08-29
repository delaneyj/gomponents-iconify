package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1V3Zm2 2v14h18V5H3Zm4 3v3h2V8h2v8H9v-3H7v3H5V8h2Zm6 0h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-4V8Zm2 2v4h2v-4h-2Z"/>`),
		g.Group(children),
	)
}