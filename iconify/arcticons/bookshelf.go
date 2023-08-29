package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookshelf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.997 10.658L6.241 19.027l17.756 5.588l17.762-8.058l-17.762-5.899z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.241 25.39l17.756 5.589l17.762-8.059"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.241 31.754l17.756 5.588l17.762-8.058M24.261 15.186l4.478 1.487M6.242 19.027c-2.411.985-2.232 5.661 0 6.364"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.758 29.285c2.411-.985 2.232-5.662 0-6.364"/>`),
		g.Group(children),
	)
}