package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Intersect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M174.63 81.37a80 80 0 1 0-93.26 93.26a80 80 0 1 0 93.26-93.26ZM100.69 136L120 155.31A63.48 63.48 0 0 1 96 160a63.48 63.48 0 0 1 4.69-24Zm33.75 11.13l-25.57-25.57a64.65 64.65 0 0 1 12.69-12.69l25.57 25.57a64.65 64.65 0 0 1-12.69 12.69ZM155.31 120L136 100.69A63.48 63.48 0 0 1 160 96a63.48 63.48 0 0 1-4.69 24ZM32 96a64 64 0 0 1 126-16a80.08 80.08 0 0 0-77.95 78A64.11 64.11 0 0 1 32 96Zm128 128a64.11 64.11 0 0 1-62-48a80.08 80.08 0 0 0 78-78a64 64 0 0 1-16 126Z"/>`),
		g.Group(children),
	)
}