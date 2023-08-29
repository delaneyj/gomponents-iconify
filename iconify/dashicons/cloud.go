package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.9 9c0-.3.1-.6.1-1c0-2.2-1.8-4-4-4c-1.6 0-2.9.9-3.6 2.2c-.2-.1-.6-.2-.9-.2C5.1 6 4 7.1 4 8.5c0 .2 0 .4.1.5c-1.8.3-3.1 1.7-3.1 3.5C1 14.4 2.6 16 4.5 16h10c1.9 0 3.5-1.6 3.5-3.5c0-1.8-1.3-3.3-3.1-3.5z"/>`),
		g.Group(children),
	)
}