package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentPasteOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 18.15l-2-2V5h-2v3h-6.15l-5-5h3.325q.275-.875 1.075-1.438T12 1q1 0 1.788.563T14.85 3H21v15.15ZM12 5q.425 0 .713-.288T13 4q0-.425-.288-.713T12 3q-.425 0-.713.288T11 4q0 .425.288.713T12 5Zm4.15 14L5 7.85V19h11.15ZM3 21V5.85L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425L18.15 21H3Z"/>`),
		g.Group(children),
	)
}