package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalPizzaOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22L2 7q2.125-1.8 4.663-2.9T12 3q2.8 0 5.338 1.088T22 7L12 22Zm0-3.6l7.3-10.95q-1.625-1.125-3.475-1.788T12 5q-1.975 0-3.813.663T4.7 7.45L12 18.4ZM9.5 10q.625 0 1.063-.438T11 8.5q0-.625-.438-1.063T9.5 7q-.625 0-1.063.438T8 8.5q0 .625.438 1.063T9.5 10Zm2.5 5q.625 0 1.063-.438T13.5 13.5q0-.625-.438-1.063T12 12q-.625 0-1.063.438T10.5 13.5q0 .625.438 1.063T12 15Zm0 3.4Z"/>`),
		g.Group(children),
	)
}