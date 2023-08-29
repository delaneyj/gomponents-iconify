package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m13 109.21l112 64a6 6 0 0 0 6 0l112-64a6 6 0 0 0 0-10.42l-112-64a6 6 0 0 0-6 0l-112 64a6 6 0 0 0 0 10.42Zm115-62.3L227.91 104L128 161.09L28.09 104ZM245.21 141a6 6 0 0 1-2.23 8.19l-112 64a6 6 0 0 1-6 0l-112-64a6 6 0 0 1 6-10.42l109 62.3l109-62.3a6 6 0 0 1 8.23 2.23Z"/>`),
		g.Group(children),
	)
}