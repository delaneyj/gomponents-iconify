package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><circle cx="4.727" cy="4.757" r="4.727"/><circle cx="15.255" cy="4.757" r="4.727"/></g>`),
		g.Group(children),
	)
}