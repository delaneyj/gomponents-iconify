package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slippers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 29h40v6H4v-6Zm19.53-16c-3.5 4-3.5 12-3.5 16h20v-8.5S28 15 23.53 13Z"/>`),
		g.Group(children),
	)
}