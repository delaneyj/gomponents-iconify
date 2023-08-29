package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TempPreferencesEcoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 6.75v-2.5q0-1.05.725-1.775T13.5 1.75H16v2.5q0 1.05-.725 1.775T13.5 6.75H11Zm-5.5 6q-1.45 0-2.475-1.025T2 9.25v-3.5h3.5q1.45 0 2.475 1.025T9 9.25v3.5H5.5Zm10.5 9q-.975 0-1.863-.3t-1.612-.825l-.825.825q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l.825-.825q-.525-.725-.825-1.613T10 15.75q0-2.5 1.75-4.263T16 9.75h5.975v6q.05 2.5-1.713 4.25T16 21.75Zm0-2q1.675 0 2.85-1.175t1.125-2.825v-4H16q-1.65-.025-2.825 1.15T12 15.75q0 .575.138 1.088t.412.962l2.75-2.75q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-2.75 2.75q.45.275.963.413T16 19.75Zm.025-4.025Z"/>`),
		g.Group(children),
	)
}