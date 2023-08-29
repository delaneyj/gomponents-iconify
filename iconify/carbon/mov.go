package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mov(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m28 9l-2 13l-2-13h-2l2.52 14h2.96L30 9h-2zM18 23h-4a2 2 0 0 1-2-2V11a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2zm-4-12v10h4V11zM8 9l-1.51 5L6 15.98L5.54 14L4 9H2v14h2v-8l-.16-2l.58 2L6 19.63L7.58 15l.58-2L8 15v8h2V9H8z"/>`),
		g.Group(children),
	)
}