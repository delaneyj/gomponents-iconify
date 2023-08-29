package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondSuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#45413c" d="M12.5 45.5a11.5 1.5 0 1 0 23 0a11.5 1.5 0 1 0-23 0Z" opacity=".15"/><path fill="#ff6242" d="M23.2 42.1L9 23.26a1 1 0 0 1 0-1.2L23.2 3.21a1 1 0 0 1 1.6 0L39 22.06a1 1 0 0 1 0 1.2L24.8 42.1a1 1 0 0 1-1.6 0Z"/><path fill="#ff866e" d="M10.39 25.16L23.2 8.21a1 1 0 0 1 1.6 0l12.81 17L39 23.26a1 1 0 0 0 0-1.2L24.8 3.21a1 1 0 0 0-1.6 0L9 22.06a1 1 0 0 0 0 1.2Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M23.2 42.1L9 23.26a1 1 0 0 1 0-1.2L23.2 3.21a1 1 0 0 1 1.6 0L39 22.06a1 1 0 0 1 0 1.2L24.8 42.1a1 1 0 0 1-1.6 0Z"/>`),
		g.Group(children),
	)
}