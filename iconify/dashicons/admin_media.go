package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdminMedia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 11V4c0-.55-.45-1-1-1h-1.67L9 1H5L3.67 3H2c-.55 0-1 .45-1 1v7c0 .55.45 1 1 1h10c.55 0 1-.45 1-1zM7 4.5a2.5 2.5 0 0 1 0 5a2.5 2.5 0 0 1 0-5zM14 6h5v10.5a2.5 2.5 0 0 1-5 0a2.5 2.5 0 0 1 3-2.45V9h-3V6zm-4 8.05V13h2v3.5a2.5 2.5 0 0 1-5 0a2.5 2.5 0 0 1 3-2.45z"/>`),
		g.Group(children),
	)
}