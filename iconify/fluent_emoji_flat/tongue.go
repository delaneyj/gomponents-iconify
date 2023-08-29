package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tongue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#89029C" d="M30 7c0 7.73-6.27 14-14 14S2 14.73 2 7h28Z"/><path fill="#F92F60" d="M19.313 12.063h-6.5c-1.079 0-4.797.593-4.797 2.906v7.969C8.016 25.343 10.062 30 16 30c5.938 0 8-4.656 8-7.063V14.97c0-2.156-3.156-2.906-4.688-2.906Z"/></g>`),
		g.Group(children),
	)
}