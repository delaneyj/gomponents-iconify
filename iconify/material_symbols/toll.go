package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 20q-3.35 0-5.675-2.325T7 12q0-3.35 2.325-5.675T15 4q3.35 0 5.675 2.325T23 12q0 3.35-2.325 5.675T15 20Zm-8-.25q-2.65-.7-4.325-2.85T1 12q0-2.75 1.675-4.9T7 4.25v2.1q-1.8.625-2.9 2.175T3 12q0 1.925 1.1 3.475T7 17.65v2.1Z"/>`),
		g.Group(children),
	)
}