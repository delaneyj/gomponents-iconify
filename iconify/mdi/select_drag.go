package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectDrag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 17h4v-4h2v4h4v2h-4v4h-2v-4h-4v-2m-2 0v2H9v-2h2m-4 0v2H5v-2h2m12-8v2h-2V9h2m0-4v2h-2V5h2m-4 0v2h-2V5h2m-4 0v2H9V5h2M7 5v2H5V5h2m0 8v2H5v-2h2m0-4v2H5V9h2Z"/>`),
		g.Group(children),
	)
}