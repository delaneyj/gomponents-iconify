package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func World(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c-4.971 0-9 4.029-9 9s4.029 9 9 9s9-4.029 9-9s-4.029-9-9-9zm2 2c0 1-.5 2-1.5 2S11 7 11 8v3s1 0 1-3a1 1 0 1 1 2 0v3a1 1 0 1 0 1 1h1v-2l1 1l-1 1c0 3 0 3-2 4c0-1-1-1-3-1v-2l-2-2V9c-1 0-1 1-1 1l-.561-.561l-2.39-2.39c.11-.192.225-.382.35-.564l.523-.678A7.977 7.977 0 0 1 12 3a8.05 8.05 0 0 1 2 .262V4z"/>`),
		g.Group(children),
	)
}