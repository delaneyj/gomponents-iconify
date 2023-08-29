package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 28a1 1 0 0 1-.71-.29l-6.13-6.14l-3.33 5a1 1 0 0 1-1 .44a1 1 0 0 1-.81-.7l-6-20A1 1 0 0 1 6.29 5l20 6a1 1 0 0 1 .7.81a1 1 0 0 1-.44 1l-5 3.33l6.14 6.13a1 1 0 0 1 0 1.42l-4 4A1 1 0 0 1 23 28Zm0-2.41L25.59 23l-7.16-7.15l5.25-3.5L7.49 7.49l4.86 16.19l3.5-5.25Z"/>`),
		g.Group(children),
	)
}