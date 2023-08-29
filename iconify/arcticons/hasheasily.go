package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hasheasily(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" d="m19.568 4.5l-4.875 39m18.614-39l-4.875 39M9.818 31.09H36.41M11.59 16.023h26.592"/>`),
		g.Group(children),
	)
}