package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DifferenceOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 11h2V9h2V7h-2V5h-2v2h-2v2h2v2Zm-2 4h6v-2h-6v2ZM6 19V1h9l6 6v12H6Zm2-2h11V8l-5-5H8v14Zm-6 6V7h2v14h11v2H2Zm6-6V3v14Z"/>`),
		g.Group(children),
	)
}