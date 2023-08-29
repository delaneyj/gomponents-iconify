package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6.537 6.516a6 6 0 0 0 7.932 7.954m2.246-1.76A6 6 0 0 0 8.3 4.277"/><path d="M12 3c1.333.333 2 2.333 2 6c0 .348 0 .681-.018 1m-.545 3.43c-.332.89-.811 1.414-1.437 1.57m0-12c-.938.234-1.547 1.295-1.825 3.182m-.156 3.837c.117 3.02.777 4.68 1.981 4.981M6 9h3m4 0h5M3 19h7m4 0h5m-9 0a2 2 0 1 0 4 0a2 2 0 1 0-4 0m2-4v2M3 3l18 18"/></g>`),
		g.Group(children),
	)
}