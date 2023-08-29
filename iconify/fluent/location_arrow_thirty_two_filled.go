package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationArrowThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M26.277 3.138c1.618-.622 3.207.967 2.585 2.585l-8.46 21.994c-.696 1.812-3.306 1.668-3.8-.209l-2.375-9.023a1 1 0 0 0-.712-.712l-9.023-2.375c-1.877-.494-2.02-3.104-.209-3.8l21.995-8.46Z"/>`),
		g.Group(children),
	)
}