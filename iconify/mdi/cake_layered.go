package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeLayered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 21v-4a2 2 0 0 0-2-2h-1v-3a2 2 0 0 0-2-2h-3V8h-2v2H8c-1.11 0-2 .89-2 2v3H5c-1.11 0-2 .89-2 2v4H1v2h22v-2M12 7a2 2 0 0 0 2-2c0-.38-.1-.73-.29-1.03L12 1l-1.72 2.97c-.18.3-.28.65-.28 1.03a2 2 0 0 0 2 2Z"/>`),
		g.Group(children),
	)
}