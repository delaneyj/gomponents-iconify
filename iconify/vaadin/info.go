package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Info(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 5h4v11H6V5zm4-3a2 2 0 1 1-3.999.001A2 2 0 0 1 10 2z"/>`),
		g.Group(children),
	)
}