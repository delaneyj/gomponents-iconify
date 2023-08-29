package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Enter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 8H1V6h5V4l3 3l-3 3zm10-8v13l-6 3v-3H4V9h1v3h5V3l4-2H5v4H4V0z"/>`),
		g.Group(children),
	)
}