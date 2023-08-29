package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongRightC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m18.73 7.757l4.25 4.236l-4.236 4.25l-1.416-1.412l1.814-1.82l-12.288.042a3.001 3.001 0 1 1-.009-2l12.316-.043l-1.843-1.836l1.412-1.417ZM5.02 12.06a1 1 0 1 0-2 .009a1 1 0 0 0 2-.01Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}