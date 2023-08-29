package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SizeWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m15.46 7l-4-2.74l-.71 1l3.09 2.11H2.16l3.09-2.1l-.71-1L.54 7a1.25 1.25 0 0 0 0 2l4 2.74l.71-1l-3.09-2.12h11.68l-3.09 2.11l.71 1l4-2.74a1.25 1.25 0 0 0 0-1.99z"/>`),
		g.Group(children),
	)
}