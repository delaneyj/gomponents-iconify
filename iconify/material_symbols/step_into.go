package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepInto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-1.25 0-2.125-.875T9 19q0-1.25.875-2.125T12 16q1.25 0 2.125.875T15 19q0 1.25-.875 2.125T12 22Zm0-8L7 9l1.4-1.4l2.6 2.575V2h2v8.175L15.575 7.6L17 9l-5 5Z"/>`),
		g.Group(children),
	)
}