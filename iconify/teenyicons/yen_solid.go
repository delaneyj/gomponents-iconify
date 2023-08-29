package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YenSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.528 7H3v1h4v3H3v1h4v3h1v-3h4v-1H8V8h4V7H8.472L12.907.79l-.814-.58L7.5 6.64L2.907.21l-.814.58L6.528 7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}