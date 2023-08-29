package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v2h8V1H3V0H0zm0 3v4.5c0 .28.22.5.5.5h7c.28 0 .5-.22.5-.5V3H0z"/>`),
		g.Group(children),
	)
}