package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trackball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M18.94 9.41a1.59 1.59 0 1 1-3.18 0a1.59 1.59 0 0 1 3.18 0Z"/><path d="M6 6.69A5.69 5.69 0 0 1 11.69 1h8.62A5.69 5.69 0 0 1 26 6.69v18.62A5.69 5.69 0 0 1 20.31 31h-8.62A5.69 5.69 0 0 1 6 25.31V6.69ZM11.69 3A3.69 3.69 0 0 0 8 6.69v18.62A3.69 3.69 0 0 0 11.69 29h8.62A3.69 3.69 0 0 0 24 25.31V6.69A3.69 3.69 0 0 0 20.31 3H17v2.083a6.002 6.002 0 0 1 0 11.834V18a1 1 0 1 1-2 0v-1.083a6.002 6.002 0 0 1 0-11.834V3h-3.31ZM16 16a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></g>`),
		g.Group(children),
	)
}