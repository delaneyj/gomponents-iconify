package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 6L7 9h2v8h2V9h2l-3-3zm8-2c0-.553-.048-1-.6-1H2.6c-.552 0-.6.447-.6 1c0 .553.048 1 .6 1h14.8c.552 0 .6-.447.6-1z"/>`),
		g.Group(children),
	)
}