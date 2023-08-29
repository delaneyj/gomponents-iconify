package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path d="M16.997 3.5v5.149c0 1.75-.964 2.425-4.595 3.358c-3.63.932-4.706 1.482-5.554 3.093C6.283 16.172 6 17.83 6 20.072V37.5"/><path stroke-linejoin="round" d="M34.942 21.509c.237 2.876-.25 5.389-1.463 7.537c-1.212 2.148-3.353 3.457-6.422 3.926M13.059 21.509c-.239 2.876.25 5.389 1.469 7.537c1.218 2.148 3.376 3.457 6.474 3.926"/><path d="M13 43.512c1.333-1.555 2-3.246 2-5.072v-8.364m20 13.436c-1.333-1.555-2-3.246-2-5.072v-8.364M31 3.5v5.149c0 1.75.964 2.425 4.595 3.358c3.63.932 4.706 1.482 5.554 3.093c.565 1.073.848 2.73.848 4.972V37.5"/></g>`),
		g.Group(children),
	)
}