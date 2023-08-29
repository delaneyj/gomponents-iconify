package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AllOutOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-5h2v3h3v2H3Zm13 0v-2h3v-3h2v5h-5Zm-4-2q-2.9 0-4.95-2.05T5 12q0-2.9 2.05-4.95T12 5q2.9 0 4.95 2.05T19 12q0 2.9-2.05 4.95T12 19Zm0-2q2.075 0 3.538-1.463T17 12q0-2.075-1.463-3.538T12 7Q9.925 7 8.462 8.463T7 12q0 2.075 1.463 3.538T12 17ZM3 8V3h5v2H5v3H3Zm16 0V5h-3V3h5v5h-2Zm-7 4Z"/>`),
		g.Group(children),
	)
}