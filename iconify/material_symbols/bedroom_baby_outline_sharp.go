package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedroomBabyOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18q2 0 3.8-.75T19 15.1l-1.05-1.05q-.275.275-.537.487t-.563.413L16 13.5V11h1v-1h-5.6L9.65 7H6l1 .75L5.5 9.5l.95 1L8 9.5v4l-.85 1.45q-.3-.2-.563-.413t-.537-.487L5 15.1q1.4 1.4 3.2 2.15T12 18Zm0-1.5q-.95 0-1.838-.188T8.45 15.7l.85-1.45q.65.25 1.338.363t1.362.112q.7 0 1.375-.112t1.325-.363l.85 1.45q-.825.375-1.713.588T12 16.5ZM2 22V2h20v20H2Zm2-2h16V4H4v16Zm0 0V4v16Z"/>`),
		g.Group(children),
	)
}