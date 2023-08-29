package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleCheckmarkSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 14h3a2.5 2.5 0 0 0 2.5-2.5v-3H8.5V14Zm0-6.5H14v-3A2.5 2.5 0 0 0 11.5 2h-3v5.5ZM7.5 2v5.5H2v-3A2.5 2.5 0 0 1 4.5 2h3ZM2 11.5v-3h5.5V14h-3A2.5 2.5 0 0 1 2 11.5Zm10.854-.646l-1.5 1.5a.5.5 0 0 1-.708 0l-.75-.75a.5.5 0 0 1 .708-.708l.396.397l1.146-1.147a.5.5 0 0 1 .708.708Z"/>`),
		g.Group(children),
	)
}