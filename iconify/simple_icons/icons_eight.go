package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IconsEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0H0v24h12zm6 12a6 6 0 1 0 0-12a6 6 0 0 0 0 12zm0 12a6 6 0 1 0 0-12a6 6 0 0 0 0 12z"/>`),
		g.Group(children),
	)
}