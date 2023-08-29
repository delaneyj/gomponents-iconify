package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionLoopRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 20h2v8h-2z"/><path fill="currentColor" d="M9 4a7.008 7.008 0 0 1 7 7v3h-2v-3a5 5 0 1 0-5 5h17.172l-4.586-4.586L23 10l7 7l-7 7l-1.414-1.414L26.172 18H9A7 7 0 0 1 9 4Z"/>`),
		g.Group(children),
	)
}