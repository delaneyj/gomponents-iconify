package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VuejsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.001 3h4l7 12l7-12h4l-11 19l-11-19Zm8.667 0L12 7l2.333-4h4.035L12 14L5.633 3h4.035Z"/>`),
		g.Group(children),
	)
}