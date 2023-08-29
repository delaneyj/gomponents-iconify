package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailAltTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 1.1L4 5.9c-1.1.4-2 1.8-2 3v8.7c0 1.2.9 1.8 2 1.4l12-4.8c1.1-.4 2-1.8 2-3V2.5c0-1.2-.9-1.8-2-1.4zm.6 2.6l-6 9.3l-6.7-4.5c-.1-.1-.4-.4-.2-.7c.2-.4.7-.2.7-.2l6.3 2.3s4.8-6.3 5.1-6.7c.1-.2.4-.3.7-.1c.3.2.2.5.1.6z"/>`),
		g.Group(children),
	)
}