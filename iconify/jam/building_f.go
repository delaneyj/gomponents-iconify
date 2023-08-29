package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20v-3a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v3H0V3a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v17H9zM4 4a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2H4zm0 3a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2H4zm0 3a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2H4zm0 3a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2H4zm4-9a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2H8zm0 3a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2H8zm0 3a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2H8zm0 3a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2H8z"/>`),
		g.Group(children),
	)
}