package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PermCameraMicRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V7q0-.825.588-1.413T4 5h3.15L8.4 3.65q.275-.3.663-.475T9.874 3h4.25q.425 0 .813.175t.662.475L16.85 5H20q.825 0 1.413.588T22 7v12q0 .825-.588 1.413T20 21h-7v-3.1q1.95-.3 3.3-1.675t1.6-3.1q.075-.45-.225-.787T16.9 12q-.35 0-.625.25t-.375.6q-.275 1.325-1.362 2.238T12 16q-1.45 0-2.538-.9T8.1 12.85q-.1-.35-.375-.6T7.1 12q-.475 0-.775.338t-.225.787q.25 1.725 1.6 3.1T11 17.9V21H4Zm8-7q.825 0 1.413-.588T14 12V8q0-.825-.588-1.413T12 6q-.825 0-1.413.588T10 8v4q0 .825.588 1.413T12 14Z"/>`),
		g.Group(children),
	)
}