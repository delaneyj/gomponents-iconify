package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Registration(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M28 25h-8a2.003 2.003 0 0 1-2-2v-3h2v3h8V9h-8v3h-2V9a2.002 2.002 0 0 1 2-2h8a2.002 2.002 0 0 1 2 2v14a2.003 2.003 0 0 1-2 2z" fill="currentColor"/><path d="M8 15h4v2H8z" fill="currentColor"/><path d="M20 15h4v2h-4z" fill="currentColor"/><path d="M14 15h4v2h-4z" fill="currentColor"/><path d="M12 25H4a2.002 2.002 0 0 1-2-2V9a2.002 2.002 0 0 1 2-2h8a2.002 2.002 0 0 1 2 2v3h-2V9H4v14h8v-3h2v3a2.002 2.002 0 0 1-2 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}