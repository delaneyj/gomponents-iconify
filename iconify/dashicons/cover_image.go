package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoverImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.2 1h15.5c.7 0 1.3.6 1.3 1.2v11.5c0 .7-.6 1.2-1.2 1.2H2.2c-.6.1-1.2-.5-1.2-1.1V2.2C1 1.6 1.6 1 2.2 1zM17 13V3H3v10h14zm-4-4s0-5 3-5v7c0 .6-.4 1-1 1H5c-.6 0-1-.4-1-1V7c2 0 3 4 3 4s1-4 3-4s3 2 3 2zm-9 8h12v2H4z"/>`),
		g.Group(children),
	)
}