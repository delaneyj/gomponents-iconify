package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedScrollingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M18 30H4a2 2 0 0 1-2-2V14a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2zM4 14v14h14V14z" fill="currentColor"/><path d="M25 23h-2V9H9V7h14a2 2 0 0 1 2 2z" fill="currentColor"/><path d="M30 16h-2V4H16V2h12a2 2 0 0 1 2 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}