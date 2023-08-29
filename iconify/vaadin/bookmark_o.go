package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 0v16l5-5l5 5V0H3zm9 13.7L8 9.8l-4 3.9V3h8v10.7zM12 2H4V1h8v1z"/>`),
		g.Group(children),
	)
}