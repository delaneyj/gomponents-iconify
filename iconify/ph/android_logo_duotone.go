package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AndroidLogoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M232 168v24a8 8 0 0 1-8 8H32a8 8 0 0 1-8-8v-22.87C24 111.65 70.15 64.2 127.63 64A104 104 0 0 1 232 168Z" opacity=".2"/><path d="M176 156a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm-84-12a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm148 24v24a16 16 0 0 1-16 16H32a16 16 0 0 1-16-16v-22.87a113.38 113.38 0 0 1 35.4-82.41L26.34 61.66a8 8 0 0 1 11.32-11.32L63.82 76.5a111.4 111.4 0 0 1 128.55-.18l26-26a8 8 0 0 1 11.32 11.32L204.82 86.5c.75.71 1.5 1.43 2.24 2.17A111.25 111.25 0 0 1 240 168Zm-16 0a96 96 0 0 0-96-96h-.34C74.91 72.18 32 115.75 32 169.13V192h192Z"/></g>`),
		g.Group(children),
	)
}