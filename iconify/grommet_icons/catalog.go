package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Catalog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M5 6L1 4.5v13.943L12 23l11-4.557V4l-4 2M5 16V2l7 3l7-3v14l-7 3l-7-3Zm6.95-11v14"/>`),
		g.Group(children),
	)
}