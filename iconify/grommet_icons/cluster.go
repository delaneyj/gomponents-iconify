package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cluster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M8 9h8V1H8v8ZM1 23h8v-8H1v8Zm14 0h8v-8h-8v8ZM5 15l3-6l-3 6Zm5 4h4h-4Zm6-10l3 6l-3-6Z"/>`),
		g.Group(children),
	)
}