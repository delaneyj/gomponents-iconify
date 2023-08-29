package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Command(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 9H5.5C3.5 9 2 7.5 2 5.5S3.5 2 5.5 2S9 3.5 9 5.5v13c0 2-1.5 3.5-3.5 3.5S2 20.5 2 18.5S3.5 15 5.5 15h13c2 0 3.5 1.5 3.5 3.5S20.5 22 18.5 22S15 20.5 15 18.5v-13c0-2 1.5-3.5 3.5-3.5S22 3.5 22 5.5S20.5 9 18.5 9H12Z"/>`),
		g.Group(children),
	)
}