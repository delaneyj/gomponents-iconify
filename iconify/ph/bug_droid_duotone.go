package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugDroidDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 120v24a80 80 0 0 1-160 0v-24Z" opacity=".2"/><path d="m191.83 43.48l13.83-13.82a8 8 0 1 0-11.32-11.32L179.79 32.9a87.81 87.81 0 0 0-103.58 0L61.66 18.34a8 8 0 0 0-11.32 11.32l13.83 13.82A87.72 87.72 0 0 0 40 104v40a88 88 0 0 0 176 0v-40a87.72 87.72 0 0 0-24.17-60.52ZM128 32a72.08 72.08 0 0 1 72 72v8H56v-8a72.08 72.08 0 0 1 72-72Zm0 184a72.08 72.08 0 0 1-72-72v-16h144v16a72.08 72.08 0 0 1-72 72Zm16-132a12 12 0 1 1 12 12a12 12 0 0 1-12-12Zm-56 0a12 12 0 1 1 12 12a12 12 0 0 1-12-12Z"/></g>`),
		g.Group(children),
	)
}