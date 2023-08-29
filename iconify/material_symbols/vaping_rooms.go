package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VapingRooms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 19q-.975-.25-1.988-.375T3 18.5H2v-2h1q1 0 2.013-.125T7 16v3Zm1 0v-3h14v3H8Zm2.5-1q.2 0 .35-.15t.15-.35q0-.2-.15-.35T10.5 17q-.2 0-.35.15t-.15.35q0 .2.15.35t.35.15Zm7.5-3v-1.3q0-.975-.575-1.488T16.05 11.7H14.5q-1.4 0-2.375-.975T11.15 8.35q0-1.4.975-2.375T14.5 5v1.5q-.75 0-1.3.525t-.55 1.325q0 .8.55 1.325t1.3.525h1.55q1.4 0 2.425.9t1.025 2.25V15H18Zm2.5 0v-2.25q0-1.65-1.15-2.85T16.5 8.7V7.2q.75 0 1.3-.55t.55-1.3q0-.75-.55-1.3t-1.3-.55V2q1.4 0 2.375.975t.975 2.375q0 .725-.275 1.313T18.85 7.75q1.4.65 2.275 2t.875 3V15h-1.5Z"/>`),
		g.Group(children),
	)
}