package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashAuto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22v-8H2V2h10l-2 7h4L5 22Zm9.625-11L18 2h1.6l3.425 9h-1.55l-.8-2.3h-3.7l-.8 2.3h-1.55Zm2.8-3.6h2.75L18.85 3.65h-.05L17.425 7.4Z"/>`),
		g.Group(children),
	)
}