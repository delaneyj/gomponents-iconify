package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDoubleUpTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.207 15.267a.75.75 0 0 1 .026-1.06l5.25-5.001a.75.75 0 0 1 1.034 0l5.25 5a.75.75 0 0 1-1.034 1.087L10 10.784l-4.733 4.51a.75.75 0 0 1-1.06-.027Zm0-4.998a.75.75 0 0 1 .026-1.06l5.25-5.002a.75.75 0 0 1 1.034 0l5.25 5.001a.75.75 0 0 1-1.034 1.086L10 5.786l-4.733 4.508a.75.75 0 0 1-1.06-.025Z"/>`),
		g.Group(children),
	)
}