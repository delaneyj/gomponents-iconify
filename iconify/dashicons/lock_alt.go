package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M15 6v2c1.1 0 2 .9 2 2v7c0 1.1-.9 2-2 2H5c-1.1 0-2-.9-2-2v-7c0-1.1.9-2 2-2V6c0-2.76 2.24-5 5-5s5 2.24 5 5zM8 6v2h4V6c0-1.1-.9-2-2-2s-2 .9-2 2zm2.8 10.25v-2.17c.7-.31 1.2-1.01 1.2-1.83c0-1.1-.9-2-2-2s-2 .9-2 2c0 .82.5 1.52 1.2 1.83v2.17h1.6z" fill="currentColor"/>`),
		g.Group(children),
	)
}