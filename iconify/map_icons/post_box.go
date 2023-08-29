package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PostBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M28.999 4.518C21.701 4.518 16 10.297 16 17.427V48h5V38h16v10h5V17.427c0-7.13-5.7-12.909-13.001-12.909zM4 2l7.364 12h.768C12.598 9 16 4.896 21 3.043V2H4z"/>`),
		g.Group(children),
	)
}