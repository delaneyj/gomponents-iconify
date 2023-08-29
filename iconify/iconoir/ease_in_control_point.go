package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EaseInControlPoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 20c8 0 18-16 18-16m-4 16a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm0 0h-2m-3 0h-2"/>`),
		g.Group(children),
	)
}