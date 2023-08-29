package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookingPotFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M88 48V16a8 8 0 0 1 16 0v32a8 8 0 0 1-16 0Zm40 8a8 8 0 0 0 8-8V16a8 8 0 0 0-16 0v32a8 8 0 0 0 8 8Zm32 0a8 8 0 0 0 8-8V16a8 8 0 0 0-16 0v32a8 8 0 0 0 8 8Zm94.4 35.2a8 8 0 0 0-11.2-1.6L224 104V88a16 16 0 0 0-16-16H48a16 16 0 0 0-16 16v16L12.8 89.6a8 8 0 0 0-9.6 12.8L32 124v60a32 32 0 0 0 32 32h128a32 32 0 0 0 32-32v-60l28.8-21.6a8 8 0 0 0 1.6-11.2Z"/>`),
		g.Group(children),
	)
}