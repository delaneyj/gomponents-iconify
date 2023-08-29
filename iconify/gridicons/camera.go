package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 12c0 1.7-1.3 3-3 3s-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3zm5-5v11c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V7c0-1.1.9-2 2-2V4h4v1h2l1-2h6l1 2h2c1.1 0 2 .9 2 2zM7.5 9c0-.8-.7-1.5-1.5-1.5S4.5 8.2 4.5 9s.7 1.5 1.5 1.5S7.5 9.8 7.5 9zM19 12c0-2.8-2.2-5-5-5s-5 2.2-5 5s2.2 5 5 5s5-2.2 5-5z"/>`),
		g.Group(children),
	)
}