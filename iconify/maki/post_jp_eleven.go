package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PostJpEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M1 2.25a.75.75 0 0 1 .75-.75h7.5a.75.75 0 0 1 0 1.5h-7.5A.75.75 0 0 1 1 2.25zm8.25 1.749h-7.5a.75.75 0 0 0 0 1.5H4.5v3a1 1 0 0 0 2 0v-3h2.75a.75.75 0 1 0 0-1.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}