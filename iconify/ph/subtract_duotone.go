package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtractDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M232 160a72 72 0 0 1-143.6 7.6a73.37 73.37 0 0 0 7.6.4a72 72 0 0 0 72-72a73.37 73.37 0 0 0-.4-7.6A72 72 0 0 1 232 160Z" opacity=".2"/><path d="M174.63 81.35a80 80 0 1 0-93.28 93.28a80 80 0 1 0 93.28-93.28ZM32 96a64 64 0 1 1 64 64a64.07 64.07 0 0 1-64-64Zm128 128a63.81 63.81 0 0 1-62-48a80.07 80.07 0 0 0 78-78a64 64 0 0 1-16 126Z"/></g>`),
		g.Group(children),
	)
}