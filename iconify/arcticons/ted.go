package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM20.811 30.226h6.226m-6.226-12.452h6.226M20.811 24h4.059m-4.059-6.226v12.452"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.25 30.226V17.774h2.802a5.448 5.448 0 0 1 5.448 5.448v1.556a5.448 5.448 0 0 1-5.448 5.448ZM9.5 17.774h8.25m-4.125 12.452V17.774"/>`),
		g.Group(children),
	)
}