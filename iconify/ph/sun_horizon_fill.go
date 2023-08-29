package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunHorizonFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M248 160a8 8 0 0 1-8 8H16a8 8 0 0 1 0-16h40.45a73.54 73.54 0 0 1-.45-8a72 72 0 0 1 144 0a73.54 73.54 0 0 1-.45 8H240a8 8 0 0 1 8 8Zm-40 32H48a8 8 0 0 0 0 16h160a8 8 0 0 0 0-16ZM80.84 59.58a8 8 0 0 0 14.32-7.16l-8-16a8 8 0 0 0-14.32 7.16Zm-60.42 43.58l16 8a8 8 0 1 0 7.16-14.31l-16-8a8 8 0 1 0-7.16 14.31ZM216 112a8 8 0 0 0 3.57-.84l16-8a8 8 0 1 0-7.16-14.31l-16 8A8 8 0 0 0 216 112Zm-51.58-48.84a8 8 0 0 0 10.74-3.58l8-16a8 8 0 0 0-14.32-7.16l-8 16a8 8 0 0 0 3.58 10.74Z"/>`),
		g.Group(children),
	)
}