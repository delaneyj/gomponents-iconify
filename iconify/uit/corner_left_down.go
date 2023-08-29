package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 5h-9.111A4.394 4.394 0 0 0 7 9.389v8.904l-4.147-4.146a.5.5 0 0 0-.707.707l5 5A.5.5 0 0 0 7.5 20c.011 0 .02-.005.03-.006a.498.498 0 0 0 .163-.033a.5.5 0 0 0 .162-.109l4.998-4.998a.5.5 0 0 0-.707-.708L8 18.293V9.389A3.393 3.393 0 0 1 11.389 6H20.5a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}