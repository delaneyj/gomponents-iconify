package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeamedNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m17 1l-.002 13c0 1.243-1.301 3-3.748 3c-1.243 0-2.25-.653-2.25-1.875c0-1.589 1.445-2.55 3-2.55c.432 0 .754.059 1 .123V5.364L8 6.637V16h-.002c0 1.243-1.301 3-3.748 3C3.007 19 2 18.347 2 17.125c0-1.589 1.445-2.55 3-2.55c.432 0 .754.059 1 .123V3l11-2z"/>`),
		g.Group(children),
	)
}