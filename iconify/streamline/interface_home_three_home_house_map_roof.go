package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceHomeThreeHomeHouseMapRoof(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M.5 7L7 .5L13.5 7m-11 1.5v5h9v-5"/>`),
		g.Group(children),
	)
}