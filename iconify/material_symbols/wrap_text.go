package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrapText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.7 20.7L11 17l3.7-3.7l1.4 1.45L14.85 16h2.4q.725 0 1.238-.513T19 14.25q0-.725-.513-1.238T17.25 12.5H4v-2h13.25q1.575 0 2.663 1.088T21 14.25q0 1.575-1.088 2.663T17.25 18h-2.4l1.25 1.25l-1.4 1.45ZM4 18v-2h5v2H4ZM4 7V5h16v2H4Z"/>`),
		g.Group(children),
	)
}