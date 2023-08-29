package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Youtube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.7 5.3c-.2-.7-.7-1.2-1.4-1.4c-2.1-.2-4.2-.4-6.3-.3c-2.1 0-4.2.1-6.3.3c-.6.2-1.2.8-1.4 1.4a37.08 37.08 0 0 0 0 9.4c.2.7.7 1.2 1.4 1.4c2.1.2 4.2.4 6.3.3c2.1 0 4.2-.1 6.3-.3c.7-.2 1.2-.7 1.4-1.4a37.08 37.08 0 0 0 0-9.4zM8 13V7l5.2 3L8 13z"/>`),
		g.Group(children),
	)
}