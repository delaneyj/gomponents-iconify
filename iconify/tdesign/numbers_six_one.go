package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumbersSixOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.42 3.5h2.31l-3.473 6.052a5.25 5.25 0 1 1-3.573 2.209L12.42 3.5Zm-3.4 9.953a3.25 3.25 0 1 0 .354-.617l-.354.617Z"/>`),
		g.Group(children),
	)
}