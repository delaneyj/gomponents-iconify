package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplayThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M131 157.4a4 4 0 0 0-6.07 0l-48 56a4 4 0 0 0 3 6.6h96a4 4 0 0 0 3-6.6ZM88.7 212l39.3-45.85L167.3 212ZM228 64v112a20 20 0 0 1-20 20h-16a4 4 0 0 1 0-8h16a12 12 0 0 0 12-12V64a12 12 0 0 0-12-12H48a12 12 0 0 0-12 12v112a12 12 0 0 0 12 12h16a4 4 0 0 1 0 8H48a20 20 0 0 1-20-20V64a20 20 0 0 1 20-20h160a20 20 0 0 1 20 20Z"/>`),
		g.Group(children),
	)
}