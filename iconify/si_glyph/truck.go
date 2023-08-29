package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Truck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(4 3)"><ellipse cx="1.438" cy="8.491" rx="1.438" ry="1.491"/><ellipse cx="9.463" cy="8.48" rx="1.463" ry="1.48"/><path d="M4.031.062v6.885h3.464c.47-.61 1.2-1.01 2.03-1.01c.831 0 1.562.399 2.031 1.01h1.398V.062H4.031z"/></g><path d="M8 11h2.916v.875H8zm8 0h.979v1H16zm-13.064.521a2.585 2.585 0 0 1 4.042-2.132V5.98c0-.527-.403-.954-.901-.954H2.946c-.499 0-1.056.846-1.056.846s-.841 1.505-.841 2.145v3.013c0 .542.396.943.901.943h1.032a2.614 2.614 0 0 1-.046-.452zm.063-5.592l3.057.021l-.022 2.112H2.253c-.068-1.276.746-2.133.746-2.133z"/></g>`),
		g.Group(children),
	)
}