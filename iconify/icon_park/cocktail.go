package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cocktail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M35.8 13H32L21 32L9.8 13H6"/><path d="M25.7509 25.5961C31.3517 28.7466 38.446 26.7602 41.5964 21.1594C44.7469 15.5586 42.7605 8.46427 37.1597 5.31383C31.5589 2.16338 24.4646 4.14978 21.3142 9.75057"/><path stroke-linejoin="round" d="M26 44H16"/><path stroke-linejoin="round" d="M21 44L21 32"/><path stroke-linejoin="round" d="M12 16C12 16 14 14 17 14C20 14 22 17 25 17C28 17 30 16 30 16"/></g>`),
		g.Group(children),
	)
}