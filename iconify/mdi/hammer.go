package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hammer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19.63L13.43 8.2l-.71-.7l1.42-1.43L12 3.89c1.2-1.19 3.09-1.19 4.27 0l3.6 3.61l-1.42 1.41h2.84l.71.71l-3.55 3.59l-.71-.71V9.62l-1.47 1.42l-.71-.71L4.13 21.76L2 19.63Z"/>`),
		g.Group(children),
	)
}