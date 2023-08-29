package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalPostOfficeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 19v3H2V12h4V8q0-2.5 1.75-4.25T12 2h4q2.5 0 4.25 1.75T22 8v14h-2v-3h-4Zm0-2h4V8q0-1.65-1.175-2.825T16 4h-4q-1.65 0-2.825 1.175T8 8v4h8v5Zm-7-.15L14 14H4l5 2.85Zm0 1.75l-5-2.85V20h10v-4.25L9 18.6ZM4 14v6v-6Zm6-4V8h8v2h-8Z"/>`),
		g.Group(children),
	)
}