package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLeftDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 88v80a8 8 0 0 1-8 8h-88v48l-96-96l96-96v48h88a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M208 72h-80V32a8 8 0 0 0-13.66-5.66l-96 96a8 8 0 0 0 0 11.32l96 96A8 8 0 0 0 128 224v-40h80a16 16 0 0 0 16-16V88a16 16 0 0 0-16-16Zm0 96h-88a8 8 0 0 0-8 8v28.69L35.31 128L112 51.31V80a8 8 0 0 0 8 8h88Z"/></g>`),
		g.Group(children),
	)
}