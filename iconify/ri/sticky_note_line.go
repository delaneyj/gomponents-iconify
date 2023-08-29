package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNoteLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 15l-6 5.996L4.002 21A.998.998 0 0 1 3 20.007V3.993C3 3.445 3.445 3 3.993 3h16.014c.548 0 .993.456.993 1.002V15ZM19 5H5v14h8v-5a1 1 0 0 1 .883-.993L14 13l5-.001V5Zm-.829 9.999L15 15v3.169l3.171-3.17Z"/>`),
		g.Group(children),
	)
}