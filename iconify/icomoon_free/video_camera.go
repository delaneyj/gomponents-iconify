package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 4.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0zm-6 0a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0zm12 5V8c0-.55-.45-1-1-1H1c-.55 0-1 .45-1 1v5c0 .55.45 1 1 1h10c.55 0 1-.45 1-1v-1.5l4 2.5V7l-4 2.5zM10 12H2V9h8v3z"/>`),
		g.Group(children),
	)
}