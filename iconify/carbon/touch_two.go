package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TouchTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29 15h-2a11 11 0 0 0-22 0H3a13 13 0 0 1 26 0Z"/><path fill="currentColor" d="M25 28h-2V15a7 7 0 1 0-14 0v13H7V15a9 9 0 0 1 18 0Z"/><path fill="currentColor" d="M21 20H11v-5a5 5 0 0 1 10 0Zm-8-2h6v-3a3 3 0 0 0-6 0Z"/>`),
		g.Group(children),
	)
}