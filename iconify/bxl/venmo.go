package bxl

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Venmo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 3.27a7.49 7.49 0 0 1 .66 3.35c0 2.72-1.93 6.72-3.49 9.27L9.53 2.44l-6.91.65L5.79 22h7.88c3.45-4.54 7.71-11 7.71-16a7.3 7.3 0 0 0-1.06-4z"/>`),
		g.Group(children),
	)
}