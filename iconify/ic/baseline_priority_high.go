package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselinePriorityHigh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12" cy="19" r="2" fill="currentColor"/><path fill="currentColor" d="M10 3h4v12h-4z"/>`),
		g.Group(children),
	)
}