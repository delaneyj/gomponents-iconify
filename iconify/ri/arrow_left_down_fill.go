package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftDownFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.361 13.052l4.95 4.95H5.997V6.687l4.95 4.95l5.657-5.658l1.414 1.415l-5.657 5.657Z"/>`),
		g.Group(children),
	)
}