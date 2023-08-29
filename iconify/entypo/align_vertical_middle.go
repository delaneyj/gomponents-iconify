package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalMiddle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10 12l-3 3h2v4h2v-4h2l-3-3zm3-7h-2V1H9v4H7l3 3l3-3zm5 5c0-.553-.048-1-.6-1H2.6c-.552 0-.6.447-.6 1c0 .551.048 1 .6 1h14.8c.552 0 .6-.449.6-1z"/>`),
		g.Group(children),
	)
}