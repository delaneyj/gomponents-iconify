package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmergencyHomeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22.025q-.4 0-.763-.15t-.662-.425L2.55 13.425q-.275-.3-.425-.663T1.975 12q0-.4.15-.775t.425-.65l8.025-8.025q.3-.3.663-.438T12 1.975q.4 0 .775.138t.65.437l8.025 8.025q.3.275.438.65t.137.775q0 .4-.137.763t-.438.662l-8.025 8.025q-.275.275-.65.425t-.775.15Zm0-2L20.025 12L12 3.975L3.975 12L12 20.025ZM12 13q.425 0 .713-.287T13 12V8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8v4q0 .425.288.713T12 13Zm0 3q.425 0 .713-.288T13 15q0-.425-.288-.713T12 14q-.425 0-.713.288T11 15q0 .425.288.713T12 16Zm0-4Z"/>`),
		g.Group(children),
	)
}