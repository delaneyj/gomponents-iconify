package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarricadeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.493 19h11.014l-.667-3H7.16l-.667 3Zm13.063 0H21v2H3v-2h1.444L7.826 3.783A1 1 0 0 1 8.802 3h6.396a1 1 0 0 1 .976.783L19.556 19ZM7.604 14h8.792l-.89-4H8.494l-.889 4Zm1.334-6h6.124l-.666-3H9.604l-.666 3Z"/>`),
		g.Group(children),
	)
}