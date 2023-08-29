package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 7a2 2 0 0 0-1.18.39l-1.75-.8L19 5.71A2 2 0 0 0 20 6a2 2 0 1 0-2-2l-1.89.87A2 2 0 1 0 15 8.5a1.88 1.88 0 0 0 .92-.24l2.1 1A2 2 0 1 0 20 7Zm-1 6a1 1 0 0 0-1 1v.39l-1.48-1.48a2.77 2.77 0 0 0-3.93 0l-.7.7l-2.48-2.49a2.79 2.79 0 0 0-3.93 0L4 12.6V7a1 1 0 0 1 1-1h5a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5a1 1 0 0 0-1-1ZM5 20a1 1 0 0 1-1-1v-3.57l2.9-2.9a.79.79 0 0 1 1.09 0l3.17 3.17l4.29 4.3Zm13-1a.89.89 0 0 1-.18.53L13.31 15l.7-.7a.78.78 0 0 1 1.1 0L18 17.22Z"/>`),
		g.Group(children),
	)
}