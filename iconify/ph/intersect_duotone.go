package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntersectDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M168 96a72 72 0 0 1-72 72a73.37 73.37 0 0 1-7.6-.4a73.37 73.37 0 0 1-.4-7.6a72 72 0 0 1 72-72a73.37 73.37 0 0 1 7.6.4a73.37 73.37 0 0 1 .4 7.6Z" opacity=".2"/><path d="M174.63 81.37a80 80 0 1 0-93.26 93.26a80 80 0 1 0 93.26-93.26ZM32 96a64 64 0 0 1 126-16a80.08 80.08 0 0 0-77.95 78A64.11 64.11 0 0 1 32 96Zm128 0a64.07 64.07 0 0 1-64 64a64.07 64.07 0 0 1 64-64Zm0 128a64.11 64.11 0 0 1-62-48a80.08 80.08 0 0 0 78-78a64 64 0 0 1-16 126Z"/></g>`),
		g.Group(children),
	)
}