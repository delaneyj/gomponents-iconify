package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlocksAndArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M20 6H6v14h14V6Zm0 22H6v14h14V28ZM42 6H28v14h14V6ZM28 28l14 14M28 28h14h-14Zm0 0v14v-14Z"/>`),
		g.Group(children),
	)
}