package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="17.061" cy="18.766" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.963" ry="5.072" transform="rotate(-40.563 17.062 18.766)"/><ellipse cx="31.39" cy="18.765" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5.963" ry="5.072" transform="rotate(-40.563 31.39 18.765)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.248 33.972s7.812 1.875 14.425-5.78"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}