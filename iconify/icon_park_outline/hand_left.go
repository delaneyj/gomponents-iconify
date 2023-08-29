package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M44 41V19h-6v22h6Z"/><path d="M38 19C30.895 12.43 26.776 8.664 25.643 7.7c-1.7-1.446-3.62-.839-3.62 2.779s5.263 5.765 5.263 8.521c.004.017-6.757.018-20.282.003A3 3 0 0 0 4 22.002v.003a3.004 3.004 0 0 0 3.004 3.005h7.013c1.206 7.968 1.872 12.301 1.997 12.998C16.202 39.054 17.2 41 20.068 41H38V19Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}