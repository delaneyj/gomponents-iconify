package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Typing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 4H4c-1.101 0-2 .9-2 2v7c0 1.1.899 2 2 2h4l4 3v-3h4c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zM6 10.6a1.1 1.1 0 1 1 0-2.2a1.1 1.1 0 0 1 0 2.2zm4 0a1.1 1.1 0 1 1 0-2.2a1.1 1.1 0 0 1 0 2.2zm4 0a1.1 1.1 0 1 1 0-2.2a1.1 1.1 0 0 1 0 2.2z"/>`),
		g.Group(children),
	)
}