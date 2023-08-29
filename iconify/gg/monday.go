package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monday(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.779 6.14a3 3 0 0 1 4.915 3.44l-5.736 8.192a3 3 0 0 1-4.915-3.441l5.736-8.192Zm8.489.088a3 3 0 0 1 4.915 3.442l-5.736 8.191a3 3 0 0 1-4.915-3.441l5.736-8.192ZM20.5 18.86a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/>`),
		g.Group(children),
	)
}