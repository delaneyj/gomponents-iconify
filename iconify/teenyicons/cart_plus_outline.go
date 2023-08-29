package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CartPlusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5.5l.6 2m0 0l2.4 8h11v-6a2 2 0 0 0-2-2H1.1ZM8.5 4v5M6 6.5h5m1.5 8a1 1 0 1 1 0-2a1 1 0 0 1 0 2Zm-8-1a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z"/>`),
		g.Group(children),
	)
}