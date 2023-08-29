package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.414 16.432L0 15.018l7.071-7.071l6.364 6.364l4.243-4.243l-1.743-1.742l6.692-1.793l-1.793 6.692l-1.742-1.742l-5.657 5.656l-6.364-6.364l-5.657 5.657Z"/>`),
		g.Group(children),
	)
}