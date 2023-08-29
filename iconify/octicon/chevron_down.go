package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M5 11L0 6l1.5-1.5L5 8.25L8.5 4.5L10 6l-5 5z" fill="currentColor"/>`),
		g.Group(children),
	)
}