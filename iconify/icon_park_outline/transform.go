package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="13.715" cy="13.714" r="6.857"/><circle cx="34.286" cy="34.285" r="6.857"/><path d="M24.001 44c-11.046 0-20-8.954-20-20l6.667 3.333M24.001 4c11.046 0 20 8.954 20 20l-6.667-3.333"/></g>`),
		g.Group(children),
	)
}