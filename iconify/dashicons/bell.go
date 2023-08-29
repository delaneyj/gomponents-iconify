package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 18c1.1 0 2-.9 2-2H8c0 1.1.9 2 2 2zm4-8.2V7.5c0-1.8-1.2-3.4-3-3.9c.1-.2.1-.4.2-.5c-.1-.6-.6-1.1-1.2-1.1s-1.1.5-1.1 1.1c0 .2.1.4.2.5c-1.8.4-3 2-3 3.9v2.2c-.1 1.2-.9 2.3-2 2.8V15h12v-2.5c-1.2-.4-2-1.5-2.1-2.7z"/>`),
		g.Group(children),
	)
}