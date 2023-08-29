package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentDismissSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 4.5V1H5.5A1.5 1.5 0 0 0 4 2.5v2.707A5.5 5.5 0 0 1 8.663 15H12.5a1.5 1.5 0 0 0 1.5-1.5V6h-3.5A1.5 1.5 0 0 1 9 4.5Zm1 0V1.25L13.75 5H10.5a.5.5 0 0 1-.5-.5ZM2.318 7.318a4.5 4.5 0 1 0 6.364 6.364a4.5 4.5 0 0 0-6.364-6.364Zm4.95 4.95a.5.5 0 0 1-.707 0L5.5 11.208l-1.06 1.06a.5.5 0 0 1-.708-.707l1.06-1.061l-1.06-1.06a.5.5 0 0 1 .707-.708L5.5 9.792l1.06-1.06a.5.5 0 0 1 .708.707L6.208 10.5l1.06 1.06a.5.5 0 0 1 0 .708Z"/>`),
		g.Group(children),
	)
}