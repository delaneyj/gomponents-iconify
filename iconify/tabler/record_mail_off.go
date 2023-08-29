package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordMailOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12a3 3 0 1 0 6 0a3 3 0 1 0-6 0m14.569 2.557a3 3 0 1 0-4.113-4.149M7 15h8M3 3l18 18"/>`),
		g.Group(children),
	)
}