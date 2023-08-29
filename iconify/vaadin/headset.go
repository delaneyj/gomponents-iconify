package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.82 8a3.016 3.016 0 0 0-1.799-1.813L13 4.5C13 2 10.53 0 7.5 0S2 2 2 4.5v1.68A3.006 3.006 0 0 0 0 9v1a3 3 0 0 0 3 3h1V6H3V4.5C3 2.57 5 1 7.5 1S12 2.57 12 4.5V6h-1v7h1a3 3 0 0 0 3-3v1.73A3.27 3.27 0 0 1 11.73 15H10a1 1 0 0 0-1-1H8a1 1 0 0 0 0 2h3.73A4.27 4.27 0 0 0 16 11.73V8h-1.18z"/>`),
		g.Group(children),
	)
}