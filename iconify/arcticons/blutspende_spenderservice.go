package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlutspendeSpenderservice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.5 17.5h-9a2 2 0 0 1-2-2v-9a1 1 0 0 0-1-1h-11a1 1 0 0 0-1 1v9a2 2 0 0 1-2 2h-9a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h9a2 2 0 0 1 2 2v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a2 2 0 0 1 2-2h9a1 1 0 0 0 1-1v-11a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}