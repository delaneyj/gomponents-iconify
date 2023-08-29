package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedScrollingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M8 30H2a2 2 0 0 1-2-2V14a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2zM2 14v14h6V14z" fill="currentColor"/><path d="M20 30h-6a2 2 0 0 1-2-2V14a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2zm-6-16v14h6V14z" fill="currentColor"/><path d="M27 21h-2V9h-8V7h8a2 2 0 0 1 2 2z" fill="currentColor"/><path d="M32 16h-2V4h-8V2h8a2 2 0 0 1 2 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}