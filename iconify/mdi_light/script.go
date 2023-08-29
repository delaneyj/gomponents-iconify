package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Script(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19a2 2 0 0 0 2 2h6.76c-.47-.53-.76-1.23-.76-2H3m11 2a2 2 0 0 0 2-2V6c0-.77.29-1.47.76-2H8a2 2 0 0 0-2 2v12h6v1a2 2 0 0 0 2 2M5 6a3 3 0 0 1 3-3h11a3 3 0 0 1 3 3v2h-5v11a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3v-1h3V6m16 1V6a2 2 0 0 0-2-2a2 2 0 0 0-2 2v1h4Z"/>`),
		g.Group(children),
	)
}