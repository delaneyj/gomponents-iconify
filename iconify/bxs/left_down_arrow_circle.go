package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftDownArrowCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.071 19.071c3.898-3.899 3.898-10.244 0-14.143c-3.899-3.899-10.244-3.898-14.143 0c-3.898 3.899-3.899 10.243 0 14.143c3.9 3.899 10.244 3.899 14.143 0zM8.464 8.464l2.829 2.829l3.535-3.536l1.414 1.414l-3.535 3.536l2.828 2.829H8.464V8.464z"/>`),
		g.Group(children),
	)
}