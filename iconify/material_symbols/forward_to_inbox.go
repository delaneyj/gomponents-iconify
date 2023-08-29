package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForwardToInbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19 23l-1.4-1.4l1.575-1.6H15v-2h4.175l-1.6-1.6L19 15l4 4l-4 4ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v7.8q-.675-.4-1.438-.6T19 13q-2.5 0-4.25 1.75T13 19v1H4Zm8-7l8-5V6l-8 5l-8-5v2l8 5Z"/>`),
		g.Group(children),
	)
}