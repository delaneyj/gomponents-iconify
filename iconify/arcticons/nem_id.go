package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NemId(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.972 23.514a6.263 6.263 0 1 1 .84 3.131m-.84-3.131H5.796c-1.279 0-2.296.706-2.296 1.697v5.54m6.488-7.237V30.7"/>`),
		g.Group(children),
	)
}