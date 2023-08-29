package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiTwoBarRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.3 14.275q-.45-.45-.425-1.075t.5-.975q1.375-1.05 3.063-1.637T12 10q1.9 0 3.588.6t3.062 1.675q.475.35.488.963T18.7 14.3q-.425.425-1.038.438t-1.137-.338q-.975-.675-2.112-1.038T12 13q-1.275 0-2.413.375T7.5 14.4q-.55.375-1.163.338T5.3 14.275ZM12 20q-.85 0-1.425-.575T10 18q0-.85.575-1.425T12 16q.85 0 1.425.575T14 18q0 .85-.575 1.425T12 20Z"/>`),
		g.Group(children),
	)
}