package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBulletListRotateNinetyTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM10 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm6.5-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM5 5.5a.5.5 0 0 0-1 0v12a.5.5 0 0 0 1 0v-12Zm5-.5a.5.5 0 0 1 .5.5v12a.5.5 0 0 1-1 0v-12A.5.5 0 0 1 10 5Zm6 .5a.5.5 0 0 0-1 0v12a.5.5 0 0 0 1 0v-12Z"/>`),
		g.Group(children),
	)
}