package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OvalFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 24c0-8.284 6.716-15 15-15h10c8.284 0 15 6.716 15 15c0 8.284-6.716 15-15 15H19c-8.284 0-15-6.716-15-15Zm15-12.5c-6.904 0-12.5 5.596-12.5 12.5S12.096 36.5 19 36.5h10c6.904 0 12.5-5.596 12.5-12.5S35.904 11.5 29 11.5H19Z"/>`),
		g.Group(children),
	)
}