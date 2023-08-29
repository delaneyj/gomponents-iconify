package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Enable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-3.3 1.925-5.938T9 2.45V4.6q-2.275.925-3.638 2.938T4 12q0 3.35 2.325 5.675T12 20q3.35 0 5.675-2.325T20 12q0-2.45-1.363-4.463T15 4.6V2.45q3.15.975 5.075 3.613T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm0-6l-5-5l1.4-1.4l2.6 2.575V2h2v10.175L15.6 9.6L17 11l-5 5Z"/>`),
		g.Group(children),
	)
}