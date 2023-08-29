package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.5 22v-7.961a20 20 0 0 0-1.567-7.761L1.5 5.25V5h6m13 17v-7.961a20 20 0 0 1 1.567-7.761L22.5 5.25V5h-6m0 17v-6.5h-9V22M24 23.5H0m12-21V5m0 0v2.5M12 5h2.5M12 5H9.5M12 15.5V22M7.5.5h9v9h-9v-9Z"/>`),
		g.Group(children),
	)
}