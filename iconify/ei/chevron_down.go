package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="m25 32.4l-9.7-9.7l1.4-1.4l8.3 8.3l8.3-8.3l1.4 1.4z"/>`),
		g.Group(children),
	)
}