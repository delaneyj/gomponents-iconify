package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextFirstLineTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.72 2.22a.75.75 0 1 1 1.06 1.06l-.97.97l.97.97a.75.75 0 0 1-1.06 1.06l-1.5-1.5a.75.75 0 0 1 0-1.06l1.5-1.5ZM3.75 3.5h7.5a.75.75 0 0 1 0 1.5h-7.5a.75.75 0 0 1 0-1.5Zm12.5 10a.75.75 0 0 1 0 1.5H3.75a.75.75 0 0 1 0-1.5h12.5ZM3.75 10h12.5a.75.75 0 0 0 0-1.5H3.75a.75.75 0 0 0 0 1.5Z"/>`),
		g.Group(children),
	)
}