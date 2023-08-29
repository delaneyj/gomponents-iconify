package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatHeartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.455 19L2 22.5V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H6.455ZM4 18.385L5.763 17H20V5H4v13.385Zm8.018-3.685l-3.359-3.36a2.25 2.25 0 0 1 3.182-3.182l.177.177l.176-.177a2.25 2.25 0 0 1 3.182 3.182l-3.358 3.36Z"/>`),
		g.Group(children),
	)
}