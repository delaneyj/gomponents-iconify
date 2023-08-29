package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.1 4h9.8a4.963 4.963 0 0 0-.9-2H3a4.963 4.963 0 0 0-.9 2zm9.8 2H2.1a5.002 5.002 0 0 0 9.8 0zM6 18v-6.07A7.002 7.002 0 0 1 2.101 0h9.798A7.002 7.002 0 0 1 8 11.93V18h2a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2h2z"/>`),
		g.Group(children),
	)
}