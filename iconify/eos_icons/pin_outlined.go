package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.675 6.83l3.546 3.537a5.004 5.004 0 0 0 2.836 1.415l-4.255 4.244a4.941 4.941 0 0 0-1.418-2.83L6.838 9.66l2.837-2.83m.709-3.537l-7.091 7.075a1.002 1.002 0 0 0 1.418 1.415l.709-.708l3.546 3.537a2.993 2.993 0 0 1 0 4.245l1.418 1.415l4.234-4.223L19.582 21H21v-1.415l-4.964-4.952l4.276-4.266l-1.418-1.415a3.01 3.01 0 0 1-4.255 0l-3.546-3.538l.71-.707a1.002 1.002 0 1 0-1.419-1.415Z"/>`),
		g.Group(children),
	)
}