package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.5 2.75l-.71.71l4.37 4.38L4 10v9a3 3 0 0 0 3 3h9c1.16 0 2.15-.66 2.64-1.62l1.82 1.83l.79-.71L7.58 7.83l-.7-.7L2.5 2.75M11 3L7.58 6.42l.71.71L11.41 4H16a2 2 0 0 1 2 2v10.84l1 1V6a3 3 0 0 0-3-3h-5M6.86 8.55l11.03 11.08C17.62 20.42 16.88 21 16 21H7a2 2 0 0 1-2-2v-8.59l1.86-1.86Z"/>`),
		g.Group(children),
	)
}