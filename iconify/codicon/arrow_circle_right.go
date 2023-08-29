package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="m8.08 5.369l2.14 2.14H4.468v1h5.68L8.08 10.574l.707.707l2.956-2.957v-.707L8.788 4.662l-.707.707Z"/><path d="M8 14A6 6 0 1 1 8 2a6 6 0 0 1 0 12Zm0-1A5 5 0 1 0 8 3a5 5 0 0 0 0 10Z"/></g>`),
		g.Group(children),
	)
}