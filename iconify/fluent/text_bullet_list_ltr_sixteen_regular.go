package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListLtrSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM2 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm1 3.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM5.5 3a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1h-9ZM5 8a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9A.5.5 0 0 1 5 8Zm.5 4a.5.5 0 0 0 0 1h9a.5.5 0 0 0 0-1h-9Z"/>`),
		g.Group(children),
	)
}