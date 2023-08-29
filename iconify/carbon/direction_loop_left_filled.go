package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionLoopLeftFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" d="M16 18h2v8h-2ZM4 15l7 7l1.414-1.414L7.828 16H21a5 5 0 1 0-5-5v1h2v-1a3 3 0 1 1 3 3H7.828l4.586-4.586L11 8Z"/><path fill="currentColor" d="M2 4v24a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Zm14 14h2v8h-2ZM4 15l7-7l1.414 1.414L7.828 14H21a3 3 0 1 0-3-3v1h-2v-1a5 5 0 1 1 5 5H7.828l4.586 4.586L11 22Z"/>`),
		g.Group(children),
	)
}