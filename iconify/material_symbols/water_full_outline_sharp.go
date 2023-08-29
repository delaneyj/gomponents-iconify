package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterFullOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.1 9q-1.375 0-2.7.388T5.95 10.55L7 20h10l1.1-10h-.7q-.95 0-1.725-.138t-2.125-.537q-.575-.175-1.2-.25T11.1 9Zm-5.4-.625Q6.975 7.7 8.338 7.35T11.125 7q.75 0 1.488.1t1.462.3q1.25.35 1.913.475T17.4 8h.925l.425-4H5.25l.45 4.375ZM5.2 22L3 2h18l-2.2 20H5.2Zm5.9-2H17H7h4.1Z"/>`),
		g.Group(children),
	)
}