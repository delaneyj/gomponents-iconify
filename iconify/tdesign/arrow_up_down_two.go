package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpDownTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 2.5v8.585l3-3L21.914 9.5L16.5 14.914L11.086 9.5L12.5 8.086l3 3V2.5h2Zm-10 6.586l5.414 5.414l-1.414 1.414l-3-3V21.5h-2v-8.586l-3 3L2.086 14.5L7.5 9.086Z"/>`),
		g.Group(children),
	)
}