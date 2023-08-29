package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnergyProgramTimeUsedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.825 0-1.412-.587Q1 20.825 1 20V6q0-.825.588-1.412Q2.175 4 3 4h8v2H3v14h14v-6h2v6q0 .825-.587 1.413Q17.825 22 17 22Zm2-4v-7h2v7Zm4 0V8h2v10Zm4 0v-4h2v4Zm5-6q-.8 0-1.512-.225q-.713-.225-1.313-.65l-.375.35q-.3.3-.7.3q-.4 0-.7-.3q-.3-.3-.3-.7q0-.4.3-.7l.4-.4q-.375-.575-.587-1.25Q13 7.75 13 7q0-2.15 1.488-3.575Q15.975 2 18 2h5v5q0 2-1.425 3.5T18 12Zm0-2q1.25 0 2.125-.875T21 7V4h-3q-1.2 0-2.1.85Q15 5.7 15 7q0 .325.062.637q.063.313.188.588l2.6-2.6q.3-.3.7-.3q.4 0 .7.3q.3.3.3.713q0 .412-.3.712l-2.6 2.6q.3.15.638.25q.337.1.712.1Z"/>`),
		g.Group(children),
	)
}