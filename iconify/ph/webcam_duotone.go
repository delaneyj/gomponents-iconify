package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebcamDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M128 32a72 72 0 1 0 72 72a72 72 0 0 0-72-72Zm0 104a32 32 0 1 1 32-32a32 32 0 0 1-32 32Z" opacity=".2"/><path d="M168 104a40 40 0 1 0-40 40a40 40 0 0 0 40-40Zm-64 0a24 24 0 1 1 24 24a24 24 0 0 1-24-24Zm120 96h-88v-16.4a80 80 0 1 0-16 0V200H32a8 8 0 0 0 0 16h192a8 8 0 0 0 0-16ZM64 104a64 64 0 1 1 64 64a64.07 64.07 0 0 1-64-64Z"/></g>`),
		g.Group(children),
	)
}