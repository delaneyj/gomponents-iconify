package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RMobiledataRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 7.35V9q0 .425-.288.713T3 10q-.425 0-.713-.288T2 9V3q0-.425.288-.713T3 2h4q.825 0 1.413.588T9 4v1.35q0 .6-.35 1.088T7.8 7.2l.65 1.525q.2.45-.075.863T7.6 10q-.275 0-.5-.15t-.35-.4l-.9-2.1H4Zm0-2h3V4H4v1.35Z"/>`),
		g.Group(children),
	)
}