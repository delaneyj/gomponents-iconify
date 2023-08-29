package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.064 3.809c-1.332-2.41-4.796-2.41-6.128 0L2.443 22.808C1.155 25.14 2.843 28 5.507 28h20.987c2.664 0 4.352-2.86 3.063-5.192l-10.493-19ZM17.25 22a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0ZM16 9a1 1 0 0 1 1 1v8a1 1 0 1 1-2 0v-8a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}