package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListLtrNinetyTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 3.25a1.25 1.25 0 1 0 2.5 0a1.25 1.25 0 0 0-2.5 0ZM15.25 7a.75.75 0 0 1 .75.75v9.5a.75.75 0 0 1-1.5 0v-9.5a.75.75 0 0 1 .75-.75ZM11 7.75a.75.75 0 0 0-1.5 0v9.5a.75.75 0 0 0 1.5 0v-9.5Zm-5 0a.75.75 0 0 0-1.5 0v9.5a.75.75 0 0 0 1.5 0v-9.5Zm4.25-3.25a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5ZM4 3.25a1.25 1.25 0 1 0 2.5 0a1.25 1.25 0 0 0-2.5 0Z"/>`),
		g.Group(children),
	)
}