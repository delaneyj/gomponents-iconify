package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ErrorTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1l11.951 20.7H.05L12 1ZM3.513 19.7h16.974L12 5L3.513 19.7ZM13 9.5V15h-2V9.5h2Zm-2 7h2.004v2.004H11V16.5Z"/>`),
		g.Group(children),
	)
}