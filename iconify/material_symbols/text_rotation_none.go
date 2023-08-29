package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotationNone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19v-2h13.2l-1.05-1.05l1.4-1.4L21 18l-3.45 3.45l-1.4-1.4L17.2 19H4Zm2.9-5L11 3h2l4.1 11h-1.9l-.95-2.8H9.8l-1 2.8H6.9Zm3.45-4.4h3.3l-1.6-4.55h-.1l-1.6 4.55Z"/>`),
		g.Group(children),
	)
}