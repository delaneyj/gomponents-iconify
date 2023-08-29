package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v16h-4v3h-2v-3H7v3H5v-3H1V3Zm2 2v12h18V5H3Zm12 4a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0ZM5 8h4v2H5V8Zm0 4h4v2H5v-2Z"/>`),
		g.Group(children),
	)
}