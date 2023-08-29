package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalfTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 1.1a.79.79 0 0 0-.717.446L4.172 3.798l-2.486.361a.8.8 0 0 0-.444 1.365l1.8 1.753l-.425 2.476a.8.8 0 0 0 1.16.844L6 9.428V1.1Z"/>`),
		g.Group(children),
	)
}