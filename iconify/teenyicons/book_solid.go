package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.5 0a.5.5 0 0 0-.5.5v13a.5.5 0 0 0 .5.5H3V0H1.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M4 15h1v-1h6.5a2.5 2.5 0 0 0 2.5-2.5v-9A2.5 2.5 0 0 0 11.5 0H4v15Zm7-10H7V4h4v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}