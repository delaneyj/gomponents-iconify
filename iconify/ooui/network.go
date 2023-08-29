package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Network(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<circle cx="10" cy="15" r="2" fill="currentColor"/><path fill="currentColor" d="M1 7.4a12 13 0 0 1 18 0l-1.5 1.4a10 11.1 0 0 0-15 0zm3.7 3.2a7 7.3 0 0 1 10.7 0L14 12a5 5.3 0 0 0-7.8 0z"/>`),
		g.Group(children),
	)
}