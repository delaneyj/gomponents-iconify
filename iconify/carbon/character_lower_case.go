package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CharacterLowerCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29 23h-5a2.003 2.003 0 0 1-2-2v-6a2.002 2.002 0 0 1 2-2h5v2h-5v6h5zM18 13h-4V9h-2v14h6a2.003 2.003 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2zm-4 8v-6h4v6zm-6-8H3v2h5v2H4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h6v-8a2.002 2.002 0 0 0-2-2zm0 8H4v-2h4z"/>`),
		g.Group(children),
	)
}