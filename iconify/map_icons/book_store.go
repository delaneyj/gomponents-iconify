package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M41 40V6c0-2.2-1.8-4-4-4H13c-2.2 0-4 1.8-4 4v38c0 2.2 1.8 4 4 4h24c1.858 0 4 0 4-2v-1H14c-1.1 0-2-.9-2-2v-3h29zM14 10c0-.55.45-1 1-1h20c.55 0 1 .45 1 1v2c0 .55-.45 1-1 1H15c-.55 0-1-.45-1-1v-2zm0 8c0-.55.45-1 1-1h20c.55 0 1 .45 1 1v2c0 .55-.45 1-1 1H15c-.55 0-1-.45-1-1v-2z"/>`),
		g.Group(children),
	)
}