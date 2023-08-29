package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartBarHorizontalOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 9H4m9 0V4.6c0-.56 0-.84-.11-1.054a.997.997 0 0 0-.435-.437C12.24 3 11.96 3 11.4 3H4v6m9 0h5.4c.56 0 .84 0 1.055.109a1 1 0 0 1 .436.437C20 9.76 20 10.04 20 10.6v2.8c0 .56 0 .84-.11 1.054a.998.998 0 0 1-.435.437C19.24 15 18.96 15 18.402 15H16M4 9v6m0 0v6h10.402c.559 0 .84 0 1.053-.109a.998.998 0 0 0 .436-.437C16 20.24 16 19.96 16 19.4V15M4 15h12"/>`),
		g.Group(children),
	)
}