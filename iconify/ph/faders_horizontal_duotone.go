package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FadersHorizontalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 80v96H40V80Z" opacity=".2"/><path d="M32 80a8 8 0 0 1 8-8h32a8 8 0 0 1 0 16H40a8 8 0 0 1-8-8Zm184 88h-40v-16a8 8 0 0 0-16 0v48a8 8 0 0 0 16 0v-16h40a8 8 0 0 0 0-16Zm-80 0H40a8 8 0 0 0 0 16h96a8 8 0 0 0 0-16Zm-32-56a8 8 0 0 0 8-8V88h104a8 8 0 0 0 0-16H112V56a8 8 0 0 0-16 0v48a8 8 0 0 0 8 8Z"/></g>`),
		g.Group(children),
	)
}