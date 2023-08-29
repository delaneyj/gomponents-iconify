package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieClapper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v4h4.865L5.532 4H4Zm4.135 0l3.333 4h4.397l-3.333-4H8.135Zm7 0l3.333 4H20V4h-4.865ZM20 10H4v10h16V10Zm-5 4H9v-2h6v2Z"/>`),
		g.Group(children),
	)
}