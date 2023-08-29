package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndeterminateQuestionBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19v-4h2v4h4v2H5Zm14 0h-4v-2h4v-4h2v4q0 .825-.588 1.413T19 21ZM3 5q0-.825.588-1.413T5 3h4v2H5v4H3V5Zm18 0v4h-2V5h-4V3h4q.825 0 1.413.588T21 5Zm-9 13q.525 0 .888-.363t.362-.887q0-.525-.363-.888T12 15.5q-.525 0-.888.363t-.362.887q0 .525.363.888T12 18Zm-.9-3.825h1.825q0-.85.2-1.3T14 11.75q.875-.875 1.163-1.413t.287-1.287q0-1.35-.975-2.2T12 6q-1.25 0-2.15.65T8.55 8.5l1.65.675q.175-.65.663-1.063T12 7.7q.725 0 1.163.388t.437 1.037q0 .5-.237.938t-.813.912q-.825.725-1.137 1.4t-.313 1.8Z"/>`),
		g.Group(children),
	)
}