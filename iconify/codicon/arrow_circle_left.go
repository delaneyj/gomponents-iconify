package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="m7.92 10.631l-2.14-2.14h5.752v-1h-5.68L7.92 5.426l-.707-.707l-2.956 2.957v.707l2.956 2.956l.707-.707Z"/><path d="M8 2a6 6 0 1 1 0 12A6 6 0 0 1 8 2Zm0 1a5 5 0 1 0 0 10A5 5 0 0 0 8 3Z"/></g>`),
		g.Group(children),
	)
}