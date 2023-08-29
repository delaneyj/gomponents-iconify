package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SupabaseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.084 15.25c-1.664 0-2.6-1.912-1.58-3.226L10.21.807C10.794.054 12 .467 12 1.42v7.33h8.916c1.663 0 2.6 1.913 1.58 3.227L13.79 23.194c-.584.753-1.79.34-1.79-.613v-7.33H3.084Z"/>`),
		g.Group(children),
	)
}