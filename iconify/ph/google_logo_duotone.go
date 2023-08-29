package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleLogoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 128a88 88 0 1 1-88-88a88 88 0 0 1 88 88Z" opacity=".2"/><path d="M224 128a96 96 0 1 1-21.95-61.09a8 8 0 1 1-12.33 10.18A80 80 0 1 0 207.6 136H128a8 8 0 0 1 0-16h88a8 8 0 0 1 8 8Z"/></g>`),
		g.Group(children),
	)
}