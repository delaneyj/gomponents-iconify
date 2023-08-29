package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IptvPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-31 13.123v10.754"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.657 29.377V18.623h3.521c1.99 0 3.604 1.617 3.604 3.611s-1.613 3.612-3.604 3.612h-3.52m8.785-7.223h7.125m-3.562 10.754V18.623m13.494 0l-3.562 10.754l-3.563-10.754"/>`),
		g.Group(children),
	)
}