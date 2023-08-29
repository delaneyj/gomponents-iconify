package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attraction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m13 2l1.5.5L13 3V2Zm1 16l1 5h-2l1-5Zm0-1.5l2 6.5h-4l2-6.5ZM12 6l9 5v2H3v-2l9-5Zm-7.5 7h15c0 4.167 1.5 10 1.5 10H3s1.5-5.833 1.5-10Zm0 0h15c0 4.167 1.5 10 1.5 10H3s1.5-5.833 1.5-10Z"/>`),
		g.Group(children),
	)
}