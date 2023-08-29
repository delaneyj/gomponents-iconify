package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m6 10l3 3v-2h8V9H9V7l-3 3zM4 2c-.553 0-1 .047-1 .6v14.8c0 .551.447.6 1 .6c.552 0 1-.049 1-.6V2.6c0-.553-.448-.6-1-.6z"/>`),
		g.Group(children),
	)
}