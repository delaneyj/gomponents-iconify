package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiberPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 15h1.5v-2H9q.425 0 .713-.288T10 12v-2q0-.425-.288-.713T9 9H5v6Zm6.25 0h1.5V9h-1.5v6ZM14 15h1.25v-3.5L17.8 15H19V9h-1.25v3.5L15.25 9H14v6Zm-7.5-3.5v-1h2v1h-2ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Z"/>`),
		g.Group(children),
	)
}