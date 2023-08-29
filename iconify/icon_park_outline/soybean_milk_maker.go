package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoybeanMilkMaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M35 10h5.064a1 1 0 0 1 .998.934l.867 13A1 1 0 0 1 40.93 25H32M7 5h28l-4 26H15l-3-21l-5-5Zm8 26h16l4 12H11l4-12Zm6 6h4"/>`),
		g.Group(children),
	)
}