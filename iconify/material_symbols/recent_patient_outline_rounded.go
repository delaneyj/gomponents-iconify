package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecentPatientOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 20H6q-.825 0-1.413-.588T4 18v-.8q0-.85.438-1.563T5.6 14.55q1.55-.775 3.15-1.163T12 13q.5 0 1 .038t1 .112v2.025q-.5-.1-1-.138T12 15q-1.4 0-2.775.338T6.5 16.35q-.225.125-.363.35T6 17.2v.8h8v2Zm-8-2h8h-8Zm6-6q-1.65 0-2.825-1.175T8 8q0-1.65 1.175-2.825T12 4q1.65 0 2.825 1.175T16 8q0 1.65-1.175 2.825T12 12Zm0-2q.825 0 1.413-.588T14 8q0-.825-.588-1.413T12 6q-.825 0-1.413.588T10 8q0 .825.588 1.413T12 10Zm0-2Zm6 11h-1q-.425 0-.713-.288T16 18v-4q0-.425.288-.713T17 13h3.775q.425 0 .65.35t.025.725L20 17h.7q.425 0 .637.375t.013.75l-2.875 5.05q-.1.175-.288.125T18 23.05V19Z"/>`),
		g.Group(children),
	)
}