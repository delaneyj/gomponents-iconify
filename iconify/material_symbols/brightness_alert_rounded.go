package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessAlertRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17q.425 0 .713-.288T13 16q0-.425-.288-.713T12 15q-.425 0-.713.288T11 16q0 .425.288.713T12 17Zm0-4q.425 0 .713-.288T13 12V8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8v4q0 .425.288.713T12 13Zm-.7 9.6L8.65 20H5q-.425 0-.713-.288T4 19v-3.65L1.4 12.7q-.275-.3-.275-.7t.275-.7L4 8.65V5q0-.425.288-.713T5 4h3.65l2.65-2.6q.3-.275.7-.275t.7.275L15.35 4H19q.425 0 .713.288T20 5v3.65l2.6 2.65q.275.3.275.7t-.275.7L20 15.35V19q0 .425-.288.713T19 20h-3.65l-2.65 2.6q-.3.275-.7.275t-.7-.275Z"/>`),
		g.Group(children),
	)
}