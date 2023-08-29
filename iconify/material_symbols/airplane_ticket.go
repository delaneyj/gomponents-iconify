package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaneTicket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.85 15.65l8.9-2.35q.375-.1.563-.463t.087-.737q-.1-.375-.437-.562t-.713-.088l-2.45.65l-4-3.75l-1.4.35l2.4 4.2l-2.4.6l-1.25-.95l-.95.25l1.65 2.85ZM20 20H4q-.825 0-1.413-.588T2 18v-4q.825 0 1.413-.588T4 12q0-.825-.588-1.413T2 10V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20Z"/>`),
		g.Group(children),
	)
}