package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrikethroughGaNaTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M9 20.25V13.5H6.837c-.82 1.332-2.155 2.421-4.183 2.738a1 1 0 1 1-.308-1.976c.77-.12 1.377-.392 1.858-.762H1.75a1 1 0 1 1 0-2h20.5a1 1 0 1 1 0 2H22v7a1 1 0 1 1-2 0v-7h-4.745l-.135 1.352l2.756-.344a1 1 0 0 1 .248 1.984l-4 .5a1 1 0 0 1-1.119-1.092l.24-2.4H11v6.75a1 1 0 1 1-2 0zm13-9.75V5a1 1 0 1 0-2 0v5.5h2zm-6.445 0l.44-4.4a1 1 0 1 0-1.99-.2l-.46 4.6h2.01zM11 10.5V4.75a1 1 0 1 0-2 0v5.75h2zm-3.113 0C7.967 9.969 8 9.46 8 9a1 1 0 0 0-1-1H2.5a1 1 0 1 0 0 2h3.438c-.021.165-.047.333-.08.5h2.03z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}