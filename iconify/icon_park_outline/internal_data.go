package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InternalData(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 18V9a3 3 0 0 1 3-3h30a3 3 0 0 1 3 3v9m-10 6v7m-8-16v16m-8-12v12M6 30v9a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3v-9"/>`),
		g.Group(children),
	)
}