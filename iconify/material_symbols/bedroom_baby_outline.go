package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedroomBabyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18q2 0 3.8-.75T19 15.1l-1.05-1.05q-.275.275-.537.487t-.563.413L16 13.5V11h1v-1h-5.6L9.65 7H6l1 .75L5.5 9.5l.95 1L8 9.5v4l-.85 1.45q-.3-.2-.563-.413t-.537-.487L5 15.1q1.4 1.4 3.2 2.15T12 18Zm0-1.5q-.95 0-1.838-.188T8.45 15.7l.85-1.45q.65.25 1.338.363t1.362.112q.7 0 1.375-.112t1.325-.363l.85 1.45q-.825.375-1.713.588T12 16.5ZM4 22q-.825 0-1.413-.588T2 20V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v16q0 .825-.588 1.413T20 22H4Zm0-2h16V4H4v16Zm0 0V4v16Z"/>`),
		g.Group(children),
	)
}