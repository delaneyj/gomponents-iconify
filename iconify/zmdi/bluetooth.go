package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m271 124l-92 92l92 92l-122 121h-21V267l-98 98l-30-30l119-119L0 97l30-30l98 98V3h21zM171 84v81l40-41zm40 224l-40-41v81z"/>`),
		g.Group(children),
	)
}