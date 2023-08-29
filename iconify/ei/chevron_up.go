package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M33.3 28.7L25 20.4l-8.3 8.3l-1.4-1.4l9.7-9.7l9.7 9.7z"/>`),
		g.Group(children),
	)
}