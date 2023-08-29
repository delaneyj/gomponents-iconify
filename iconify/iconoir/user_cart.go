package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m22 12.5l-.833 2.5m0 0L20 18.5h-4.5l-1-3.5h6.667ZM16.5 20.51l.01-.011m2.99.011l.01-.011M9 11a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M2 18a7 7 0 0 1 11.33-5.5"/></g>`),
		g.Group(children),
	)
}