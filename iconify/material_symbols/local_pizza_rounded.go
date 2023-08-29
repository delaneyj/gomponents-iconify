package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalPizzaRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20.4q-.475 0-.925-.225t-.75-.675l-7.2-10.825q-.475-.725-.338-1.588T3.6 5.776Q5.475 4.5 7.6 3.75T12 3q2.275 0 4.4.738t4 2.037q.65.45.8 1.313t-.325 1.587l-7.2 10.825q-.3.45-.75.675T12 20.4ZM9.5 10q.625 0 1.063-.438T11 8.5q0-.625-.438-1.063T9.5 7q-.625 0-1.063.438T8 8.5q0 .625.438 1.063T9.5 10Zm2.5 5q.625 0 1.063-.438T13.5 13.5q0-.625-.438-1.063T12 12q-.625 0-1.063.438T10.5 13.5q0 .625.438 1.063T12 15Z"/>`),
		g.Group(children),
	)
}