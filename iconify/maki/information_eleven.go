package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InformationEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.599.94c-.6 0-1.1.5-1.1 1.1s.5 1.1 1.1 1.1s1.1-.5 1.1-1.1s-.5-1.1-1.1-1.1zM3 4l-.001.74S4.5 4.634 4.5 6v1.5c0 1.5-1.501 1.74-1.501 1.74L3 10h5.2l-.001-.76s-1.2 0-1.2-1.5L7 5s0-1-1-1H3z" fill="currentColor"/>`),
		g.Group(children),
	)
}