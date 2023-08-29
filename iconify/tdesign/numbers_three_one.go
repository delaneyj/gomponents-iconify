package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumbersThreeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 4h9v3.02l-4.264 2.986A5 5 0 0 1 12 20h-.5a5 5 0 0 1-5-5h2a3 3 0 0 0 3 3h.5a3 3 0 1 0 0-6H9.5V9.48L14.47 6H7.5V4Z"/>`),
		g.Group(children),
	)
}