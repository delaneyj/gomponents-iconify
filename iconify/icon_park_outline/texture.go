package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Texture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m12 6l-6 6m36 24l-6 6M22 6L6 22M32 6L6 32M42 6L6 42m36-26L16 42m26-16L26 42"/>`),
		g.Group(children),
	)
}