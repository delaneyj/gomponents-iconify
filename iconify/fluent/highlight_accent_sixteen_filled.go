package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighlightAccentSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.5 2v3a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5V2h11Zm-8 6.5h5v1.867a2 2 0 0 1-.971 1.716L5.5 14.5v-6Z"/>`),
		g.Group(children),
	)
}