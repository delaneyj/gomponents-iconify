package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusPlaceholder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<rect width="20" height="20" x="2" y="2" fill="none" stroke="currentColor" stroke-width="2" rx="2"/>`),
		g.Group(children),
	)
}