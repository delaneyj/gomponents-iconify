package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M27.3 34.7L17.6 25l9.7-9.7l1.4 1.4l-8.3 8.3l8.3 8.3z"/>`),
		g.Group(children),
	)
}