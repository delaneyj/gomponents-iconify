package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineSegmentBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M217.47 38.53a36 36 0 0 0-57.95 41l-80 80a36.07 36.07 0 0 0-41 7a36 36 0 1 0 58 9.95l80-80a36 36 0 0 0 41-57.95Zm-145 162a12 12 0 1 1 0-17a12 12 0 0 1 .01 16.97Zm128-128a12 12 0 0 1-17 0a12 12 0 1 1 17 0Z"/>`),
		g.Group(children),
	)
}