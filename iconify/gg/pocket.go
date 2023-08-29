package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4h18v7a9 9 0 1 1-18 0V4ZM1 2h22v9c0 6.075-4.925 11-11 11S1 17.075 1 11V2Zm10.293 12.694a1 1 0 0 0 1.414 0l4.243-4.243a1 1 0 1 0-1.414-1.414L12 12.572L8.464 9.037A1 1 0 0 0 7.05 10.45l4.243 4.242Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}