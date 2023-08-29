package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterClear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.57 3h18.86l-6.93 9.817V21h-5v-8.183L2.57 3Zm3.86 2l5.07 7.183V19h1v-6.817L17.57 5H6.43Zm11.45 8.465L20 15.586l2.122-2.121l1.414 1.414L21.415 17l2.121 2.122l-1.414 1.414L20 18.414l-2.12 2.122l-1.415-1.415l2.121-2.12l-2.121-2.122l1.414-1.414Z"/>`),
		g.Group(children),
	)
}