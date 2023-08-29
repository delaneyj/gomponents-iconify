package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 11h-2V3H9v8H7l3 3l3-3zm4.4 4H2.6c-.552 0-.6.447-.6 1c0 .553.048 1 .6 1h14.8c.552 0 .6-.447.6-1c0-.553-.048-1-.6-1z"/>`),
		g.Group(children),
	)
}