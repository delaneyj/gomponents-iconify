package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPinTimeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 11V6h-2v7h6v-2h-4Zm5.364 6.364L12 23.728l-6.364-6.364a9 9 0 1 1 12.728 0Z"/>`),
		g.Group(children),
	)
}