package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DbTwoDataSharingGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 28v-6.17l2.59 2.58L12 23l-5-5l-5 5l1.41 1.41L6 21.83V28c0 1.103.897 2 2 2h9v-2H8zm17-11c-2.85 0-5 1.29-5 3v7c0 1.71 2.15 3 5 3s5-1.29 5-3v-7c0-1.71-2.15-3-5-3zm0 2c1.936 0 3 .751 3 1s-1.064 1-3 1s-3-.751-3-1s1.064-1 3-1zm0 9c-1.936 0-3-.751-3-1v-4.58c.826.363 1.85.58 3 .58s2.174-.217 3-.58V27c0 .249-1.064 1-3 1zm3.59-20.41L26 10.17V4c0-1.103-.897-2-2-2h-9v2h9v6.17l-2.59-2.58L20 9l5 5l5-5l-1.41-1.41zM7 15c2.85 0 5-1.29 5-3V5c0-1.71-2.15-3-5-3S2 3.29 2 5v7c0 1.71 2.15 3 5 3zM7 4c1.936 0 3 .751 3 1S8.936 6 7 6s-3-.751-3-1s1.064-1 3-1zM4 7.42c.826.363 1.85.58 3 .58s2.174-.217 3-.58V12c0 .249-1.064 1-3 1s-3-.751-3-1V7.42z"/>`),
		g.Group(children),
	)
}