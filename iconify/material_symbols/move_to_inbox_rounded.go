package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveToInboxRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm7-5q.8 0 1.475-.413t1.1-1.087q.15-.225.375-.363t.5-.137H19V5H5v9h3.55q.275 0 .5.138t.375.362q.425.675 1.1 1.088T12 16Zm-1-5.85V7q0-.425.288-.713T12 6q.425 0 .713.288T13 7v3.15l.875-.875q.15-.15.338-.225t.375-.063q.187.013.375.088t.337.225q.275.3.288.7t-.288.7l-2.6 2.6q-.15.15-.325.213t-.375.062q-.2 0-.375-.062T11.3 13.3l-2.6-2.6q-.15-.15-.225-.337T8.4 9.988q0-.188.075-.363T8.7 9.3q.3-.3.713-.313t.712.288l.875.875Z"/>`),
		g.Group(children),
	)
}