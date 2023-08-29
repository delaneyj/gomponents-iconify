package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrderDescending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4h13v2H2V4Zm0 7h13v2H2v-2Zm1 7H2v2h11v-2H3Zm16 3.414l.707-.707l3-3l.707-.707L22 15.586l-.707.707L20 17.586V4h-2v13.586l-1.293-1.293l-.707-.707L14.586 17l.707.707l3 3l.707.707Z"/>`),
		g.Group(children),
	)
}