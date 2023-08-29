package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerTenSelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 16h3V8h-3v8Zm0 3q-1.25 0-2.125-.875T7 16V8q0-1.25.875-2.125T10 5h3q1.25 0 2.125.875T16 8v8q0 1.25-.875 2.125T13 19h-3Zm-7 0V8H1V5h5v14H3Zm14 0v-2h4v-1h-2.65q-.575 0-.963-.388T17 14.65v-2.3q0-.575.388-.963T18.35 11H23v2h-4v1h2.65q.575 0 .963.388t.387.962v2.3q0 .575-.388.963T21.65 19H17Z"/>`),
		g.Group(children),
	)
}