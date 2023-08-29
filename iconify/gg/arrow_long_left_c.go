package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongLeftC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m5.27 7.757l-4.25 4.236l4.236 4.25l1.416-1.412l-1.814-1.82l12.288.042a3.001 3.001 0 0 0 5.834-.975a3 3 0 0 0-5.825-1.025L4.839 11.01l1.843-1.836L5.27 7.757Zm13.71 4.303a1 1 0 1 1 2 .009a1 1 0 0 1-2-.01Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}