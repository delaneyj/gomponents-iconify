package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListLtrTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.25 6a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5ZM7 4.75A.75.75 0 0 1 7.75 4h9.5a.75.75 0 0 1 0 1.5h-9.5A.75.75 0 0 1 7 4.75ZM7.75 9a.75.75 0 0 0 0 1.5h9.5a.75.75 0 0 0 0-1.5h-9.5Zm0 5a.75.75 0 0 0 0 1.5h9.5a.75.75 0 0 0 0-1.5h-9.5ZM4.5 9.75a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0ZM3.25 16a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Z"/>`),
		g.Group(children),
	)
}