package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infinity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.121 9.879l2.083 2.083l.007-.006l1.452 1.452l.006.006l2.122 2.122a5 5 0 1 0 0-7.072l-.714.714l1.415 1.414l.713-.713a3 3 0 1 1 0 4.242l-2.072-2.072l-.007.006l-3.59-3.59a5 5 0 1 0 0 7.07l.713-.713l-1.414-1.414l-.714.713a3 3 0 1 1 0-4.242Z"/>`),
		g.Group(children),
	)
}