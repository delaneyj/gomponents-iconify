package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlickrCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><circle cx="7.364" cy="10.379" r="2.364"/><circle cx="12.628" cy="10.379" r="2.364"/><path d="M10 18.494a8 8 0 1 0 0-16a8 8 0 0 0 0 16zm0 2c-5.523 0-10-4.478-10-10c0-5.523 4.477-10 10-10s10 4.477 10 10c0 5.522-4.477 10-10 10z"/></g>`),
		g.Group(children),
	)
}