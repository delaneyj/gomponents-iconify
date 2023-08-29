package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Treinta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.883 14.174h9.797c-.618 7.22-4.846 9.03-9.797 9.758V43.5a9.466 9.466 0 0 1-9.563-9.673V14.173A9.466 9.466 0 0 1 23.883 4.5v9.673m9.797 25.518a3.814 3.814 0 1 1-3.814-3.809a3.811 3.811 0 0 1 3.814 3.81Z"/>`),
		g.Group(children),
	)
}