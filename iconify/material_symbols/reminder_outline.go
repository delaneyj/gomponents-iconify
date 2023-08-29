package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReminderOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 8h1V7q0-.425-.288-.713T6.5 6q-.425 0-.713.288T5.5 7q0 .425.288.713T6.5 8ZM11 8q.425 0 .713-.288T12 7q0-.425-.288-.713T11 6q-.425 0-.713.288T10 7v1h1Zm1.55 7.5ZM10.475 22q-.7 0-1.312-.3t-1.038-.85l-5.45-6.925l.475-.5q.5-.525 1.2-.625t1.3.275L7.5 14.2V10h-1q-1.25 0-2.125-.875T3.5 7q0-1.25.875-2.125T6.5 4q.275 0 .513.05t.487.125V3q0-.425.288-.713T8.5 2q.425 0 .725.288t.3.712v1.4q.35-.2.713-.3T11 4q1.25 0 2.125.875T14 7q0 1.25-.875 2.125T11 10H9.525v7.8L7.1 16.3l2.6 3.325q.15.175.35.275t.425.1H16q.825 0 1.413-.588T18 18v-4q0-.425-.288-.713T17 13h-5.475v-2H17q1.25 0 2.125.875T20 14v4q0 1.65-1.175 2.825T16 22h-5.525Z"/>`),
		g.Group(children),
	)
}