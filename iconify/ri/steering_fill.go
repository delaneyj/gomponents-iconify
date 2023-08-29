package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SteeringFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.8 14.001a10.009 10.009 0 0 1-8.401 7.902v-2.025A8.01 8.01 0 0 0 19.747 14l2.052.001Zm-17.548 0a8.01 8.01 0 0 0 6.247 5.859v2.028a10.01 10.01 0 0 1-8.3-7.887h2.053ZM17.999 11v2h-1a4 4 0 0 0-3.995 3.8L13 17v1h-2v-1a4 4 0 0 0-3.8-3.995L7 13H6v-2h12Zm-6-9c5.186 0 9.45 3.947 9.951 9h-2.012A8.001 8.001 0 0 0 4.06 11H2.05C2.552 5.947 6.815 2 12 2Z"/>`),
		g.Group(children),
	)
}