package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayListRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.964 4.634h-12v2h12v-2Zm0 4h-12v2h12v-2Zm-12 4h8v2h-8v-2Zm9 1.076l1.415-1.415l2.121 2.121l2.121-2.12l1.415 1.413l-2.122 2.122l2.122 2.12l-1.415 1.415l-2.121-2.121l-2.121 2.121l-1.415-1.414l2.122-2.122l-2.122-2.12Z"/>`),
		g.Group(children),
	)
}