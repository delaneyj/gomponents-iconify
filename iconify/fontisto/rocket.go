package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 24h3.15c-.176-.634-.552-1.906-1.13-4.9L.001 24zm10.62-4.9c-.578 2.994-.954 4.266-1.13 4.9h3.15l-2.018-4.9zm-4.303-6.438a2.297 2.297 0 1 1 0-4.594a2.297 2.297 0 0 1 0 4.594zm4.525-2.976C10.594 2.437 6.317 0 6.317 0S2.039 2.436 1.792 9.686C1.562 16.406 4.06 24 4.06 24h4.514s2.498-7.594 2.268-14.314z"/>`),
		g.Group(children),
	)
}