package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CssThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.7 3.4l-.6 3.2h12.3L17 8.7H4.7l-.6 3.2h12.3l-.7 3.6l-5 1.7l-4.3-1.7l.3-1.6h-3L3 17.7l7.1 2.9l8.2-2.9l1.1-5.8l.2-1.2L21 3.4H5.7z"/>`),
		g.Group(children),
	)
}