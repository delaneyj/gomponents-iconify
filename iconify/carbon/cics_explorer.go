package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CicsExplorer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m24 16l-4.6-1.4l2.3-4.3l-4.3 2.3L16 8l-1.4 4.6l-4.3-2.3l2.3 4.3L8 16l4.6 1.4l-2.3 4.3l4.3-2.3L16 24l1.4-4.6l4.3 2.3l-2.3-4.3L24 16z"/><path fill="currentColor" d="M16 30a14 14 0 1 1 14-14a14.016 14.016 0 0 1-14 14Zm0-26a12 12 0 1 0 12 12A12.014 12.014 0 0 0 16 4Z"/>`),
		g.Group(children),
	)
}