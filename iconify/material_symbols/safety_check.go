package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SafetyCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-3.475-.875-5.738-3.988T4 11.1V5l8-3l8 3v6.1q0 3.8-2.263 6.913T12 22Zm0-5q2.075 0 3.538-1.463T17 12q0-2.075-1.463-3.538T12 7Q9.925 7 8.462 8.463T7 12q0 2.075 1.463 3.538T12 17Zm1.65-2.65L11.5 12.2V9h1v2.8l1.85 1.85l-.7.7Z"/>`),
		g.Group(children),
	)
}