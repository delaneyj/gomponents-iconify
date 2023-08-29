package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-1.525 0-3.125-.55t-2.9-1.5q-1.3-.95-2.138-2.225T3 15v-3l4 3l-1.55 1.55q.725 1.275 2.3 2.2T11 19.925V11H8V9h3V7.825q-.875-.325-1.438-1.088T9 5q0-1.25.875-2.125T12 2q1.25 0 2.125.875T15 5q0 .975-.563 1.738T13 7.824V9h3v2h-3v8.925q1.675-.25 3.25-1.175t2.3-2.2L17 15l4-3v3q0 1.45-.838 2.725t-2.137 2.225q-1.3.95-2.9 1.5T12 22Zm0-16q.425 0 .713-.288T13 5q0-.425-.288-.713T12 4q-.425 0-.713.288T11 5q0 .425.288.713T12 6Z"/>`),
		g.Group(children),
	)
}