package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9 4l-.89-2.66A.5.5 0 0 0 7.64 1H3.36a.5.5 0 0 0-.47.34L2 4a1 1 0 0 0-1 1v3h1v1a1 1 0 1 0 2 0V8h3v1a1 1 0 1 0 2 0V8h1V5a1 1 0 0 0-1-1zM3 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0-3l.62-2h3.76L8 4H3zm5 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}