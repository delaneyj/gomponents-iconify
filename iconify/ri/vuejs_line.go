package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VuejsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.317 3L12 18l8.684-15H23L12 22L1 3h2.316Zm4.342 0L12 10.5L16.343 3h2.316L12 14.5L5.343 3H7.66Z"/>`),
		g.Group(children),
	)
}