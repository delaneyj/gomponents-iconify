package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M198 128a6 6 0 0 1-6 6H64a6 6 0 0 1 0-12h128a6 6 0 0 1 6 6Zm34-54H24a6 6 0 0 0 0 12h208a6 6 0 0 0 0-12Zm-80 96h-48a6 6 0 0 0 0 12h48a6 6 0 0 0 0-12Z"/>`),
		g.Group(children),
	)
}