package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCreamDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 112a16 16 0 0 1-16 16H56a16 16 0 0 1 0-32v-8a72 72 0 0 1 144 0v8a16 16 0 0 1 16 16Z" opacity=".2"/><path d="M208 89.37V88a80 80 0 0 0-160 0v1.37A24 24 0 0 0 56 136h3.36l61.69 108a8 8 0 0 0 13.9 0l61.69-108H200a24 24 0 0 0 8-46.63Zm-80 134.51L77.79 136h19.28l40.57 71Zm6.79-87.88l21.71 38l-9.64 16.88L115.5 136Zm30.92 21.88L153.21 136h25ZM200 120H56a8 8 0 0 1 0-16a8 8 0 0 0 8-8v-8a64 64 0 0 1 128 0v8a8 8 0 0 0 8 8a8 8 0 0 1 0 16Z"/></g>`),
		g.Group(children),
	)
}