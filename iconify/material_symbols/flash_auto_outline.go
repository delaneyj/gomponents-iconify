package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashAutoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7 15.6l3.2-4.6H7.35l2-7H4v8h3v3.6ZM5 22v-8H2V2h10l-2 7h4L5 22Zm2-10H4h3Zm7.625-1L18 2h1.6l3.425 9h-1.55l-.8-2.3h-3.7l-.8 2.3h-1.55Zm2.8-3.6h2.75L18.85 3.65h-.05L17.425 7.4Z"/>`),
		g.Group(children),
	)
}