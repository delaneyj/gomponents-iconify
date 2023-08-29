package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PartOfCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M78.2 0v512L453 164C359.5 63.4 226.4 0 78.2 0z"/>`),
		g.Group(children),
	)
}