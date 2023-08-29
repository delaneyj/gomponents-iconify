package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 5c0-.55-.45-1-1-1H2c-.55 0-1 .45-1 1c0 .57.49 1 1 1h5c.55 0 1-.45 1-1zm6 5l4-4v10l-4-4v-2zm-1 4V8c0-.55-.45-1-1-1H4c-.55 0-1 .45-1 1v6c0 .55.45 1 1 1h8c.55 0 1-.45 1-1z"/>`),
		g.Group(children),
	)
}