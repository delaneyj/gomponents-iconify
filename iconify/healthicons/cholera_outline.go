package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CholeraOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M31.938 13A8.001 8.001 0 0 0 24 6H10V4h14c5.185 0 9.449 3.947 9.95 9H36v6H24v-6h2a1 1 0 0 0-1-1H10v-2h15a3 3 0 0 1 3 3h3.938ZM26 17v-2h8v2h-8Z" clip-rule="evenodd"/><path d="M27.03 33.517a1 1 0 0 1-1.012-.988A2.5 2.5 0 0 1 28.488 30a1 1 0 0 1 .024 2a.5.5 0 0 0-.494.506a1 1 0 0 1-.988 1.011ZM32 31a1 1 0 1 0 0 2a.5.5 0 0 1 .5.5a1 1 0 1 0 2 0A2.5 2.5 0 0 0 32 31Zm-2.237 7.25a1 1 0 0 1-.762 1.192a2.5 2.5 0 0 1-2.978-1.906a1 1 0 0 1 1.954-.429a.5.5 0 0 0 .595.381a1 1 0 0 1 1.191.763ZM33 39a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm-2-11a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-1 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path fill-rule="evenodd" d="M38 35.636C38 40.256 34.418 44 30 44s-8-3.745-8-8.364C22 28.318 30 21 30 21s8 7.318 8 14.636Zm-2 0c0 3.6-2.77 6.364-6 6.364s-6-2.764-6-6.364c0-3.004 1.681-6.229 3.616-8.884A33.405 33.405 0 0 1 30 23.85c.68.74 1.533 1.734 2.384 2.9C34.319 29.408 36 32.633 36 35.637Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}