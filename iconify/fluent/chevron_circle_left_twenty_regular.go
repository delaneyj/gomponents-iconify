package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleLeftTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 3a7 7 0 1 1 0 14a7 7 0 0 1 0-14Zm8 7a8 8 0 1 0-16 0a8 8 0 0 0 16 0Zm-6.146-3.146a.5.5 0 0 0-.708-.708l-3.5 3.5a.5.5 0 0 0 0 .708l3.5 3.5a.5.5 0 0 0 .708-.708L8.707 10l3.147-3.146Z"/>`),
		g.Group(children),
	)
}