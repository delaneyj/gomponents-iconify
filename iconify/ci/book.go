package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 19.5V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C6.52 3 7.08 3 8.2 3h9.2c.56 0 .84 0 1.055.109a1 1 0 0 1 .436.437C19 3.76 19 4.04 19 4.6v11.8c0 .56 0 .84-.11 1.054a.998.998 0 0 1-.435.437C18.24 18 17.96 18 17.402 18H7.25A2.25 2.25 0 0 0 5 20.25c0 .414.336.75.75.75h10.652c.559 0 .84 0 1.053-.109a.998.998 0 0 0 .436-.437C18 20.24 18 19.96 18 19.4V18"/>`),
		g.Group(children),
	)
}