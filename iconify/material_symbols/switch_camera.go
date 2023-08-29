package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V7q0-.825.588-1.413T4 5h3.15L9 3h6l1.85 2H20q.825 0 1.413.588T22 7v12q0 .825-.588 1.413T20 21H4Zm5-4l1.4-1.4l-1.55-1.55h6.3L13.6 15.6L15 17l4-4l-4-4l-1.4 1.4l1.6 1.65H8.8l1.6-1.65L9 9l-4 4l4 4Z"/>`),
		g.Group(children),
	)
}