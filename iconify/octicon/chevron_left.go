package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M5.5 3L7 4.5L3.25 8L7 11.5L5.5 13l-5-5l5-5z" fill="currentColor"/>`),
		g.Group(children),
	)
}