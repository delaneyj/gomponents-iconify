package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusDisabledSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<rect width="10" height="10" x="1" y="1" fill="currentColor" fill-rule="evenodd" stroke="currentColor" rx="1"/>`),
		g.Group(children),
	)
}