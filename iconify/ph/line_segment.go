package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineSegment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M214.64 41.36a32 32 0 0 0-50.2 38.89l-84.19 84.19a32.06 32.06 0 0 0-38.89 4.94a32 32 0 1 0 50.2 6.37l84.19-84.19a32 32 0 0 0 38.89-50.2Zm-139.33 162a16 16 0 0 1-22.64-22.64a16 16 0 0 1 22.63 0a16 16 0 0 1 .01 22.61Zm128-128a16 16 0 1 1 0-22.63a16 16 0 0 1 .02 22.57Z"/>`),
		g.Group(children),
	)
}