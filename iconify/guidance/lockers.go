package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lockers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor"><path d="M.5 9.5v.25l4.75 4.75h.25a3.89 3.89 0 0 1 5.5 0l9 9h3.5V20l-1-1h-2v-2h-2v-2l-4-4a3.89 3.89 0 0 1 0-5.5v-.25L9.75.5H9.5a9 9 0 0 0-9 9Z"/><path d="M4.5 6a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z"/></g>`),
		g.Group(children),
	)
}