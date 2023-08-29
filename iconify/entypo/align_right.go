package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 7v2H3v2h8v2l3-3l-3-3zm4-4.4v14.8c0 .551.448.6 1 .6c.553 0 1-.049 1-.6V2.6c0-.553-.447-.6-1-.6c-.552 0-1 .047-1 .6z"/>`),
		g.Group(children),
	)
}