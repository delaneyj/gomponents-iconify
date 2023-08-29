package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelLeftFocusRightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.004 13a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h8ZM6 4v8H4.004a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1H6Z"/>`),
		g.Group(children),
	)
}