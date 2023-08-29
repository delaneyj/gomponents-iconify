package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M196 128a4 4 0 0 1-4 4H64a4 4 0 0 1 0-8h128a4 4 0 0 1 4 4Zm36-52H24a4 4 0 0 0 0 8h208a4 4 0 0 0 0-8Zm-80 96h-48a4 4 0 0 0 0 8h48a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}