package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M16 3H8a5 5 0 0 0-5 5v2a5 5 0 0 0 4.82 4.997l-.563 3.378a2.254 2.254 0 0 0 3.817 1.965L16 15.414V3zm2 12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3v12z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}