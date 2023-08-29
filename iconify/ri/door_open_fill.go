package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorOpenFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.998 21v-2h2V4.835a1 1 0 0 1 .821-.984l9.472-1.722a.6.6 0 0 1 .707.59V4h4a1 1 0 0 1 1 1v14h2v2h-4V6h-3v15h-13Zm10-10h-2v2h2v-2Z"/>`),
		g.Group(children),
	)
}