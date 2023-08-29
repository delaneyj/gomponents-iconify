package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProductHuntFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><g clip-path="url(#akarIconsProductHuntFill0)"><path fill="currentColor" d="M15.402 10.2c0 .992-.808 1.8-1.8 1.8H10.2V8.4h3.402c.992 0 1.8.808 1.8 1.8ZM24 12c0 6.629-5.371 12-12 12S0 18.629 0 12S5.371 0 12 0s12 5.371 12 12Zm-6.198-1.8c0-2.318-1.883-4.2-4.2-4.2H7.8v12h2.4v-3.6h3.402c2.317 0 4.2-1.882 4.2-4.2Z"/></g><defs><clipPath id="akarIconsProductHuntFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}