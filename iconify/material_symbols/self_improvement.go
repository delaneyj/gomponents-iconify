package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelfImprovement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.8 20q-.75 0-1.275-.525T5 18.2q0-.525.3-.987t.8-.663L10 15v-2.25q-1.35 1.575-3.138 2.413T3 16v-2q1.7 0 3.088-.7t2.512-2l1.35-1.6q.3-.35.7-.525T11.5 9h1q.45 0 .85.175t.7.525l1.35 1.6q1.125 1.3 2.512 2T21 14v2q-2.075 0-3.863-.838T14 12.75V15l3.9 1.55q.5.2.8.663t.3.987q0 .75-.525 1.275T17.2 20H10v-.5q0-.65.425-1.075T11.5 18h3q.225 0 .363-.138T15 17.5q0-.225-.138-.363T14.5 17h-3q-1.05 0-1.775.725T9 19.5v.5H6.8ZM12 8q-.825 0-1.413-.588T10 6q0-.825.588-1.413T12 4q.825 0 1.413.588T14 6q0 .825-.588 1.413T12 8Z"/>`),
		g.Group(children),
	)
}