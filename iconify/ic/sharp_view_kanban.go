package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpViewKanban(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3H3v18h18V3zM9 17H7V7h2v10zm4-5h-2V7h2v5zm4 3h-2V7h2v8z"/>`),
		g.Group(children),
	)
}