package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalMiddle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 10L5 7v2H1v2h4v2l3-3zm7 3v-2h4V9h-4V7l-3 3l3 3zm-5 5c.553 0 1-.049 1-.6V2.6c0-.553-.447-.6-1-.6c-.552 0-1 .047-1 .6v14.8c0 .551.448.6 1 .6z"/>`),
		g.Group(children),
	)
}