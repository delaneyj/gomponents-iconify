package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuizSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 15q.425 0 .738-.313t.312-.737q0-.425-.313-.737T14 12.9q-.425 0-.738.313t-.312.737q0 .425.313.738T14 15Zm-.75-3.2h1.5q0-.725.15-1.063t.7-.887q.75-.75 1-1.212t.25-1.088q0-1.125-.788-1.837T14 5q-1.025 0-1.788.575T11.15 7.1l1.35.55q.225-.625.613-.938T14 6.4q.6 0 .975.338t.375.912q0 .35-.2.663t-.7.787q-.825.725-1.012 1.137T13.25 11.8ZM6 18V2h16v16H6Zm-4 4V6h2v14h14v2H2Z"/>`),
		g.Group(children),
	)
}