package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationNorth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.54 19.977a.34.34 0 0 0 .357-.07a.33.33 0 0 0 .084-.35L12 9L7.018 19.557a.33.33 0 0 0 .084.35a.34.34 0 0 0 .357.07L12 18.5l4.54 1.477zM12 3v2"/>`),
		g.Group(children),
	)
}