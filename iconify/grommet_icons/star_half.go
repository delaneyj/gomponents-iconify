package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path fill-opacity=".2" d="M12 16.667V2l2.5 7.5H22L16 14l3 8z"/><path d="M12 16.667L5 22l3-8l-6-4.5h7.5L12 2z"/></g>`),
		g.Group(children),
	)
}