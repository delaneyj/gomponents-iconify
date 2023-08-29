package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gateway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 20v-5m0-6V4m-7 8h5m9 0h-5m-.969-4.031L12 9.344l-1.031-1.375h2.062zm-2.062 8.07L12 14.664l1.031 1.375H10.97zM6 13.031L4.625 12L6 10.969v2.062zm12-2.062L19.375 12L18 13.031V10.97zM12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1z"/>`),
		g.Group(children),
	)
}