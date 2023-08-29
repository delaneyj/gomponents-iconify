package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineBarRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21q-.425 0-.713-.288T8 20q0-.425.288-.713T9 19h2v-4.1q-2.15-.35-3.575-2T6 9V4q0-.425.288-.713T7 3h10q.425 0 .713.288T18 4v5q0 2.25-1.425 3.9T13 14.9V19h2q.425 0 .713.288T16 20q0 .425-.288.713T15 21H9ZM8 8h8V5H8v3Z"/>`),
		g.Group(children),
	)
}