package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.662 5.521L5.237 19l.707-4.967l-4.945.709L14.424 1.263c.391-.392 1.133-.308 1.412 0l2.826 2.839c.5.473.391 1.026 0 1.419z"/>`),
		g.Group(children),
	)
}