package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactsBookFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2v20H3V2h4Zm2 0h10.005C20.107 2 21 2.898 21 3.99v16.02c0 1.099-.893 1.99-1.995 1.99H9V2Zm13 4h2v4h-2V6Zm0 6h2v4h-2v-4Zm-7 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-3 4h6a3 3 0 1 0-6 0Z"/>`),
		g.Group(children),
	)
}