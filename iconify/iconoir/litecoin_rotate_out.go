package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LitecoinRotateOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.168 8A10.003 10.003 0 0 0 12 2c-5.185 0-9.45 3.947-9.95 9"/><path stroke-linejoin="round" d="M18 8h3.4a.6.6 0 0 0 .6-.6V4M2.881 16c1.544 3.532 5.068 6 9.168 6c5.186 0 9.45-3.947 9.951-9"/><path stroke-linejoin="round" d="M6.05 16h-3.4a.6.6 0 0 0-.6.6V20"/><path d="M10.5 7v9.4a.6.6 0 0 0 .6.6h4.4"/><path stroke-linejoin="round" d="m8.5 13l4.5-2"/></g>`),
		g.Group(children),
	)
}