package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PermCameraMicOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V7q0-.825.588-1.413T4 5h3.15L8.4 3.65q.275-.3.663-.475T9.874 3h4.25q.425 0 .813.175t.662.475L16.85 5H20q.825 0 1.413.588T22 7v12q0 .825-.588 1.413T20 21h-4q-.425 0-.713-.288T15 20q0-.425.288-.713T16 19h4V7h-4.05l-1.825-2h-4.25L8.05 7H4v12h4q.425 0 .713.288T9 20q0 .425-.288.713T8 21H4Zm8-5q-1.45 0-2.538-.9T8.1 12.85q-.1-.35-.375-.6T7.1 12q-.475 0-.775.338t-.225.787Q6.5 15 7.825 16.313T11 17.9V20q0 .425.288.713T12 21q.425 0 .713-.288T13 20v-2.1q1.85-.275 3.175-1.588t1.725-3.187q.075-.45-.225-.788T16.9 12q-.35 0-.625.25t-.375.6q-.275 1.325-1.362 2.238T12 16Zm0-2q.825 0 1.413-.588T14 12V8q0-.825-.588-1.413T12 6q-.825 0-1.413.588T10 8v4q0 .825.588 1.413T12 14Zm-8 5h16H4Z"/>`),
		g.Group(children),
	)
}