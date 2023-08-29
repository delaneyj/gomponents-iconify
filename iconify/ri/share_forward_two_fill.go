package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareForwardTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19h16v-5h2v6a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1v-6h2v5Zm8-9H9a5.992 5.992 0 0 0-4.854 2.473A8.003 8.003 0 0 1 12 6V2l8 6l-8 6v-4Z"/>`),
		g.Group(children),
	)
}