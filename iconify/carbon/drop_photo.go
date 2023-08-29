package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropPhoto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 26a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4Zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2Z"/><path fill="currentColor" d="M27 29H11a2 2 0 0 1-2-2v-9a2 2 0 0 1 2-2h3.28l.543-1.632A2 2 0 0 1 16.721 13h4.558a2 2 0 0 1 1.898 1.368L23.72 16H27a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2zm-16-2h16v-9h-4.72l-1-3h-4.56l-1 3H11zm16-16h2v2h-2zm0-4h2v2h-2zm0-4h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zm-4 0h2v2h-2zM7 3h2v2H7zM3 3h2v2H3zm0 4h2v2H3zm0 4h2v2H3zm0 4h2v2H3zm0 4h2v2H3zm0 4h2v2H3zm0 4h2v2H3z"/>`),
		g.Group(children),
	)
}