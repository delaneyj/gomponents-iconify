package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bamowi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.933 5.5h6.328a.949.949 0 0 1 .949.949v2.846h4.393a1.897 1.897 0 0 1 1.897 1.897v29.41a1.897 1.897 0 0 1-1.897 1.898H23.896a1.897 1.897 0 0 1-1.898-1.897v-29.41a1.897 1.897 0 0 1 1.898-1.898h4.089V6.45a.949.949 0 0 1 .948-.949Zm-6.935 30.359H42.5M27.985 9.295h8.225"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 24h2.567l1.739-4.645l2.242 13.458l2.243-21.589l2.243 16.879L19.506 24h2.492"/>`),
		g.Group(children),
	)
}