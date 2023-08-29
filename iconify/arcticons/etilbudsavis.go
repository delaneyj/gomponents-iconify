package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Etilbudsavis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.5 10.5h-32v27a5 5 0 0 0 5 5h27a5 5 0 0 0 5-5v-22a5 5 0 0 0-5-5Zm-32 0l27-5v5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.126 31.97a6.664 6.664 0 0 1-5.793 3.363h-6.666A6.667 6.667 0 0 1 14 28.667v-4.334a6.667 6.667 0 0 1 6.667-6.666h6.666A6.667 6.667 0 0 1 34 24.333V26.5H14"/>`),
		g.Group(children),
	)
}