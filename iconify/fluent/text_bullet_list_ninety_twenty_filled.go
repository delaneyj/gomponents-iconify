package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListNinetyTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 16.75a1.25 1.25 0 1 0-2.5 0a1.25 1.25 0 0 0 2.5 0ZM4.75 13a.75.75 0 0 1-.75-.75v-9.5a.75.75 0 0 1 1.5 0v9.5a.75.75 0 0 1-.75.75ZM9 12.25a.75.75 0 0 0 1.5 0v-9.5a.75.75 0 0 0-1.5 0v9.5Zm5 0a.75.75 0 0 0 1.5 0v-9.5a.75.75 0 0 0-1.5 0v9.5ZM9.75 15.5a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5ZM16 16.75a1.25 1.25 0 1 0-2.5 0a1.25 1.25 0 0 0 2.5 0Z"/>`),
		g.Group(children),
	)
}