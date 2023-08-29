package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlusAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.8 4.2c3.2 3.21 3.2 8.39 0 11.6a8.208 8.208 0 0 1-11.6 0a8.208 8.208 0 0 1 0-11.6C7.41 1 12.59 1 15.8 4.2zm-4.3 11.3v-4h4v-3h-4v-4h-3v4h-4v3h4v4h3z"/>`),
		g.Group(children),
	)
}