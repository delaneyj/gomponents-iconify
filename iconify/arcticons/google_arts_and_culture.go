package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleArtsAndCulture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.7 15.372a5.134 5.134 0 1 1 5.134 5.134v5.738c6.005 0 10.872-4.868 10.872-10.872S27.84 4.5 21.834 4.5S10.962 9.368 10.962 15.372H16.7Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.484 37.22a8.755 8.755 0 0 1-3.648.797a8.749 8.749 0 0 1-8.758-8.758a8.756 8.756 0 0 1 8.758-8.758v-5.134c-7.67 0-13.892 6.22-13.892 13.892c0 7.682 6.222 13.892 13.892 13.892a13.82 13.82 0 0 0 7.55-2.235m.109-7.417a8.69 8.69 0 0 0 1.1-4.24v-.967h7.84v5.436h-3.456a13.483 13.483 0 0 1-1.691 3.358"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.113 43.5h7.943L21.834 26.244h-7.942L32.113 43.5zm-10.277-5.483v5.134"/>`),
		g.Group(children),
	)
}