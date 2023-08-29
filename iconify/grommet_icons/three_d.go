package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M2 5.5v13l10 4l10-4v-13l-10-4l-10 4ZM13 12h5v5M2 5.5l10 4l10-4M6 14s6 6.5 12-2"/>`),
		g.Group(children),
	)
}