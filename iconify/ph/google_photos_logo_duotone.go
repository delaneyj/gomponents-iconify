package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePhotosLogoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M192 88a63.69 63.69 0 0 1-14 40h-50V24a64 64 0 0 1 64 64ZM64 168a64 64 0 0 0 64 64V128H78a63.69 63.69 0 0 0-14 40Z" opacity=".2"/><path d="M232 120h-39.51A72 72 0 0 0 128 16a8 8 0 0 0-8 8v39.51A72 72 0 0 0 16 128a8 8 0 0 0 8 8h39.51A72 72 0 0 0 128 240a8 8 0 0 0 8-8v-39.51A72 72 0 0 0 240 128a8 8 0 0 0-8-8ZM120 223.43A56.09 56.09 0 0 1 72 168a55.31 55.31 0 0 1 10-32h38ZM120 120H32.57A56.09 56.09 0 0 1 88 72a55.31 55.31 0 0 1 32 10Zm16-87.43A56.09 56.09 0 0 1 184 88a55.31 55.31 0 0 1-10 32h-38ZM168 184a55.31 55.31 0 0 1-32-10v-38h87.43A56.09 56.09 0 0 1 168 184Z"/></g>`),
		g.Group(children),
	)
}