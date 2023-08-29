package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumbersFiveOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 4H16v2H9.5v4H12a5.002 5.002 0 0 1 4.9 4h.1v1a5 5 0 0 1-5 5h-.5a5 5 0 0 1-5-5h2a3 3 0 0 0 3 3h.5a3 3 0 1 0 0-6H7.5V4Z"/>`),
		g.Group(children),
	)
}