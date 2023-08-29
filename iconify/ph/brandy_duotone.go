package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandyDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 88a88 88 0 0 1-176 0Z" opacity=".2"/><path d="M224 88a95.63 95.63 0 0 0-15.53-52.37a8 8 0 0 0-6.7-3.63H54.23a8 8 0 0 0-6.7 3.63A95.63 95.63 0 0 0 32 88a96.12 96.12 0 0 0 88 95.66V216H88a8 8 0 0 0 0 16h80a8 8 0 0 0 0-16h-32v-32.34A96.12 96.12 0 0 0 224 88ZM58.7 48h138.6a79.52 79.52 0 0 1 10.3 32H48.4a79.52 79.52 0 0 1 10.3-32ZM128 168a80.11 80.11 0 0 1-79.6-72h159.2a80.11 80.11 0 0 1-79.6 72Z"/></g>`),
		g.Group(children),
	)
}