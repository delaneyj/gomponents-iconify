package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignDistributedVerticalTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m16.78 5.28l.72-.72v16.69a.75.75 0 0 0 1.5 0V4.56l.72.72a.75.75 0 1 0 1.06-1.06l-2-2a.75.75 0 0 0-1.06 0l-2 2a.75.75 0 0 0 1.06 1.06ZM6 19.44l.72-.72a.75.75 0 0 1 1.06 1.06l-2 2a.75.75 0 0 1-1.06 0l-2-2a.75.75 0 1 1 1.06-1.06l.72.72V2.75a.75.75 0 0 1 1.5 0v16.69Zm6.5-16.69a.75.75 0 0 0-1.5 0v18.5a.75.75 0 0 0 1.5 0V2.75Z"/>`),
		g.Group(children),
	)
}