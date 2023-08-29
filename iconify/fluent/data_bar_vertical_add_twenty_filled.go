package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataBarVerticalAddTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 3a2 2 0 0 1 2 2v4.6a5.47 5.47 0 0 0-2.5-.6a5.5 5.5 0 0 0-1.5.207V5a2 2 0 0 1 2-2ZM9 14.5c0 .882.208 1.716.577 2.455A2 2 0 0 1 8 15V8a2 2 0 1 1 4 0v1.6a5.5 5.5 0 0 0-3 4.9ZM7 11a2 2 0 1 0-4 0v4a2 2 0 1 0 4 0v-4Zm12 3.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-4-2a.5.5 0 0 0-1 0V14h-1.5a.5.5 0 0 0 0 1H14v1.5a.5.5 0 0 0 1 0V15h1.5a.5.5 0 0 0 0-1H15v-1.5Z"/>`),
		g.Group(children),
	)
}