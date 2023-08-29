package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockCountdownThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 136.33A100.13 100.13 0 1 1 119.67 28a4 4 0 1 1 .66 8A92.13 92.13 0 1 0 220 135.67a4 4 0 1 1 8 .66ZM128 132h56a4 4 0 0 0 0-8h-52V72a4 4 0 0 0-8 0v56a4 4 0 0 0 4 4Zm32-88a8 8 0 1 0-8-8a8 8 0 0 0 8 8Zm36 24a8 8 0 1 0-8-8a8 8 0 0 0 8 8Zm24 36a8 8 0 1 0-8-8a8 8 0 0 0 8 8Z"/>`),
		g.Group(children),
	)
}