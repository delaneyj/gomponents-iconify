package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23.414 22.001L2 .587L.586 2l.914.914v19.79L6.876 18h9.71L22 23.415l1.414-1.414Zm-8.828-6H6.124L3.5 18.297V4.915l11.086 11.086ZM22.503 18L22.5 2H9.655l-3.418-.003L8.24 4H20.5v11.995L22.503 18Z"/>`),
		g.Group(children),
	)
}