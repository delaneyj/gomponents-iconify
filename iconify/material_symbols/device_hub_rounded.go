package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceHubRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.425 0-.713-.288T3 20v-3q0-.425.288-.713T4 16h3l4-4V8.8q-.9-.325-1.45-1.088T9 6q0-1.25.875-2.125T12 3q1.25 0 2.125.875T15 6q0 .95-.55 1.713T13 8.8V12l4 4h3q.425 0 .713.288T21 17v3q0 .425-.288.713T20 21h-3q-.425 0-.713-.288T16 20v-2.05l-4-4l-4 4V20q0 .425-.288.713T7 21H4Z"/>`),
		g.Group(children),
	)
}