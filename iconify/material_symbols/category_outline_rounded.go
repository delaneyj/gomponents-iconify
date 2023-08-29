package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CategoryOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.425 9.475L11.15 3.4q.15-.25.375-.363T12 2.925q.25 0 .475.113t.375.362l3.725 6.075q.15.25.15.525t-.125.5q-.125.225-.35.363t-.525.137h-7.45q-.3 0-.525-.138T7.4 10.5q-.125-.225-.125-.5t.15-.525ZM17.5 22q-1.875 0-3.188-1.313T13 17.5q0-1.875 1.313-3.188T17.5 13q1.875 0 3.188 1.313T22 17.5q0 1.875-1.313 3.188T17.5 22ZM3 20.5v-6q0-.425.288-.713T4 13.5h6q.425 0 .713.288T11 14.5v6q0 .425-.288.713T10 21.5H4q-.425 0-.713-.288T3 20.5Zm14.5-.5q1.05 0 1.775-.725T20 17.5q0-1.05-.725-1.775T17.5 15q-1.05 0-1.775.725T15 17.5q0 1.05.725 1.775T17.5 20ZM5 19.5h4v-4H5v4ZM10.05 9h3.9L12 5.85L10.05 9ZM12 9Zm-3 6.5Zm8.5 2Z"/>`),
		g.Group(children),
	)
}