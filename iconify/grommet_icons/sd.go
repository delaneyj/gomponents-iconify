package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 2a2 2 0 0 1 2-2h13a1 1 0 0 1 .707.293l4 4A1 1 0 0 1 22 5v17a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V2Zm14.586 0H4v20h16V5.414L16.586 2ZM5 7V3h2v4H5Zm3-4v4h2V3H8Zm3 4V3h2v4h-2Zm3-2v2h2V5h-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}