package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextFormat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19v-2h14v2H5Zm1.9-4L11 4h2l4.1 11h-1.9l-.95-2.8H9.8l-1 2.8H6.9Zm3.45-4.4h3.3l-1.6-4.55h-.1l-1.6 4.55Z"/>`),
		g.Group(children),
	)
}