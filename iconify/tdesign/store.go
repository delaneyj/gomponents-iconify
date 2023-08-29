package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Store(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 2.5h17v2h-17v-2Zm-.82 3h18.64l1.18 5.901V13.5H21v8h-2v-8h-5v8H3v-8H1.5v-2.099L2.68 5.5Zm2.32 8v6h7v-6H5Zm15.48-2l-.8-4H4.32l-.8 4h16.96Z"/>`),
		g.Group(children),
	)
}