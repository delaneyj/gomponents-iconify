package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-1.25 0-2.125-.875T9 19q0-1.25.875-2.125T12 16q1.25 0 2.125.875T15 19q0 1.25-.875 2.125T12 22Zm-1-8V5.825L8.4 8.4L7 7l5-5l5 5l-1.425 1.4L13 5.825V14h-2Z"/>`),
		g.Group(children),
	)
}