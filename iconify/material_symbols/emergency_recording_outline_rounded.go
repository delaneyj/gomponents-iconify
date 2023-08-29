package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmergencyRecordingOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 13.725V16q0 .425.288.713T10 17q.425 0 .713-.288T11 16v-2.275l2.15 1.225q.35.2.75.1t.6-.45q.2-.35.088-.763t-.463-.612L12 12l2.125-1.225q.35-.2.463-.612T14.5 9.4q-.2-.35-.6-.45t-.75.1L11 10.275V8q0-.425-.288-.713T10 7q-.425 0-.713.288T9 8v2.275L6.85 9.05q-.35-.2-.75-.1t-.6.45q-.2.35-.088.763t.463.612L8 12l-2.125 1.225q-.35.2-.463.613t.088.762q.2.35.6.45t.75-.1L9 13.725ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h12q.825 0 1.413.588T18 6v4.5l3.15-3.15q.25-.25.55-.125t.3.475v8.6q0 .35-.3.475t-.55-.125L18 13.5V18q0 .825-.588 1.413T16 20H4Zm0-2h12V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}