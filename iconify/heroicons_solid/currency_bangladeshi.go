package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyBangladeshi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 18a8 8 0 1 0 0-16a8 8 0 0 0 0 16ZM7 4a1 1 0 0 0 0 2a1 1 0 0 1 1 1v1H7a1 1 0 0 0 0 2h1v3a3 3 0 1 0 6 0v-1a1 1 0 1 0-2 0v1a1 1 0 1 1-2 0v-3h3a1 1 0 1 0 0-2h-3V7a3 3 0 0 0-3-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}