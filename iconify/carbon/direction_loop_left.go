package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionLoopLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 20h2v8h-2z"/><path fill="currentColor" d="M23 4a7.008 7.008 0 0 0-7 7v3h2v-3a5 5 0 1 1 5 5H5.828l4.586-4.586L9 10l-7 7l7 7l1.414-1.414L5.828 18H23a7 7 0 0 0 0-14Z"/>`),
		g.Group(children),
	)
}