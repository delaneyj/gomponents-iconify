package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectUngroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M1 1h3v3H1V1Zm12 0h3v3h-3V1ZM4 2h9m2 7h5M4 15h9M1 13h3v3H1v-3Zm12 0h3v3h-3v-3ZM2 4v9m13-9v9m5-5h3v3h-3V8Zm-9 14h9M8 20h3v3H8v-3Zm12 0h3v3h-3v-3ZM9 16v4m13-9v9"/>`),
		g.Group(children),
	)
}