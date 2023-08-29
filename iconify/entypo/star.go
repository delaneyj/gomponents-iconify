package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10 1.3l2.388 6.722H18.8l-5.232 3.948l1.871 6.928L10 14.744l-5.438 4.154l1.87-6.928l-5.233-3.948h6.412L10 1.3z"/>`),
		g.Group(children),
	)
}