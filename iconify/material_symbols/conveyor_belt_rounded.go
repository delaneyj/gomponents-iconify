package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConveyorBeltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-1.25 0-2.125-.875T2 18q0-1.25.875-2.125T5 15h14q1.25 0 2.125.875T22 18q0 1.25-.875 2.125T19 21H5Zm0-2h14q.425 0 .713-.288T20 18q0-.425-.288-.713T19 17H5q-.425 0-.713.288T4 18q0 .425.288.713T5 19Zm5-6q-.425 0-.713-.288T9 12V4q0-.425.288-.713T10 3h8q.425 0 .713.288T19 4v8q0 .425-.288.713T18 13h-8Zm-7.05-2.05q-.4 0-.675-.288T2 9.976q0-.4.288-.675t.687-.275h3q.4 0 .675.275t.275.675q0 .4-.275.688t-.675.287H2.95ZM13 8h2.025q.425 0 .7-.288T16 7q0-.425-.275-.7t-.7-.275h-2.05q-.425 0-.7.275T12 7q0 .425.288.712T13 8ZM4.975 8q-.425 0-.7-.288T4 7q0-.425.288-.7T5 6.025h.95q.425 0 .7.275t.275.7q0 .425-.275.713T5.95 8h-.975Z"/>`),
		g.Group(children),
	)
}