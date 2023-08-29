package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RainbowFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 48A120.13 120.13 0 0 0 8 168v16a8 8 0 0 0 8 8h224a8 8 0 0 0 8-8v-16A120.13 120.13 0 0 0 128 48Zm32 128a8 8 0 0 1-8-8a24 24 0 0 0-48 0a8 8 0 0 1-16 0a40 40 0 0 1 80 0a8 8 0 0 1-8 8Zm32 0a8 8 0 0 1-8-8a56 56 0 0 0-112 0a8 8 0 0 1-16 0a72 72 0 0 1 144 0a8 8 0 0 1-8 8Zm32 0a8 8 0 0 1-8-8a88 88 0 0 0-176 0a8 8 0 0 1-16 0a104 104 0 0 1 208 0a8 8 0 0 1-8 8Z"/>`),
		g.Group(children),
	)
}