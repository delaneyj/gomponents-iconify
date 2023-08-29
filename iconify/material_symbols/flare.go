package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 13v-2h6v2H1Zm6.75-3.85l-2.1-2.1l1.4-1.4l2.1 2.1l-1.4 1.4ZM11 7V1h2v6h-2Zm5.25 2.15l-1.4-1.4l2.1-2.1l1.4 1.4l-2.1 2.1ZM17 13v-2h6v2h-6Zm-5 2q-1.25 0-2.125-.875T9 12q0-1.25.875-2.125T12 9q1.25 0 2.125.875T15 12q0 1.25-.875 2.125T12 15Zm4.95 3.35l-2.1-2.1l1.4-1.4l2.1 2.1l-1.4 1.4Zm-9.9 0l-1.4-1.4l2.1-2.1l1.4 1.4l-2.1 2.1ZM11 23v-6h2v6h-2Z"/>`),
		g.Group(children),
	)
}