package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastleJpEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M1.75 9.503a.75.75 0 0 1-.75-.75V4.748a.75.75 0 0 1 .75-.75l1.251-.001v-1.75a.75.75 0 0 1 .75-.75h3.5a.75.75 0 0 1 .75.75v1.75l1.25.001a.75.75 0 0 1 .75.75v4a.75.75 0 0 1-1.5 0v-3.25l-1.25-.001a.75.75 0 0 1-.75-.75v-1.75h-2v1.75a.75.75 0 0 1-.75.75L2.5 5.498v3.255a.75.75 0 0 1-.75.75z" fill="currentColor"/>`),
		g.Group(children),
	)
}