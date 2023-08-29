package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Balance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21v-2h9V7.825q-.65-.225-1.125-.7T9.175 6H6l3 7q0 1.25-1.025 2.125T5.5 16q-1.45 0-2.475-.875T2 13l3-7H3V4h6.175q.3-.875 1.075-1.438T12 2q.975 0 1.75.563T14.825 4H21v2h-2l3 7q0 1.25-1.025 2.125T18.5 16q-1.45 0-2.475-.875T15 13l3-7h-3.175q-.225.65-.7 1.125t-1.125.7V19h9v2H2Zm14.625-8h3.75L18.5 8.65L16.625 13Zm-13 0h3.75L5.5 8.65L3.625 13ZM12 6q.425 0 .713-.288T13 5q0-.425-.288-.713T12 4q-.425 0-.713.288T11 5q0 .425.288.713T12 6Z"/>`),
		g.Group(children),
	)
}