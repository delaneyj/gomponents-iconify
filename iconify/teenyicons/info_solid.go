package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 .99V2H7V.99h1ZM7 13H2v1h11v-1H8V4H4v1h3v8Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}