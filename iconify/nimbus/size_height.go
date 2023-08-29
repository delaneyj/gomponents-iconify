package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SizeHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m9 15.46l2.74-4l-1-.71l-2.1 3.09V2.16l2.1 3.09l1-.71L9 .54a1.25 1.25 0 0 0-2 0l-2.74 4l1 .71l2.12-3.09v11.68l-2.11-3.09l-1 .71l2.74 4a1.25 1.25 0 0 0 1.99 0z"/>`),
		g.Group(children),
	)
}