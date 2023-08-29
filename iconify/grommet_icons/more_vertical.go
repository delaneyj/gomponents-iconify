package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 14h-4v-4h4v4Zm0-7h-4V3h4v4Zm0 14h-4v-4h4v4Z"/>`),
		g.Group(children),
	)
}