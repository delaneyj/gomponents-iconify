package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RainbowDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M240 168v16h-64v-16a48 48 0 0 0-96 0v16H16v-16a112 112 0 0 1 224 0Z" opacity=".2"/><path d="M184 168v16a8 8 0 0 1-16 0v-16a40 40 0 0 0-80 0v16a8 8 0 0 1-16 0v-16a56 56 0 0 1 112 0Zm-56-88a88.1 88.1 0 0 0-88 88v16a8 8 0 0 0 16 0v-16a72 72 0 0 1 144 0v16a8 8 0 0 0 16 0v-16a88.1 88.1 0 0 0-88-88Zm0-32A120.13 120.13 0 0 0 8 168v16a8 8 0 0 0 16 0v-16a104 104 0 0 1 208 0v16a8 8 0 0 0 16 0v-16A120.13 120.13 0 0 0 128 48Z"/></g>`),
		g.Group(children),
	)
}