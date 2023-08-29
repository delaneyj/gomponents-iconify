package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.465 14H8a1 1 0 0 1 0-2h2v-2H8a1 1 0 1 1 0-2h2V6H8a1 1 0 1 1 0-2h1.465A3.998 3.998 0 0 0 6 2c-1.48 0-2.773.804-3.465 2H4a1 1 0 1 1 0 2H2v2h2a1 1 0 1 1 0 2H2v2h2a1 1 0 0 1 0 2H2.535A3.998 3.998 0 0 0 6 16c1.48 0 2.773-.804 3.465-2zm-1.492 3.668a2 2 0 1 1-3.945 0A6.003 6.003 0 0 1 0 12V6a6 6 0 1 1 12 0v6a6.003 6.003 0 0 1-4.027 5.668z"/>`),
		g.Group(children),
	)
}