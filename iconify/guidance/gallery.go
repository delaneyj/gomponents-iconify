package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gallery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M0 19.5c3.314 0 6-2.5 6-4.5m0 0c0 2 2.686 4.5 6 4.5s6-2.5 6-4.5M6 15v8.5M18 15c0 2 2.686 4.5 6 4.5M18 15v8.5m-14 0h4m8 0h4m1.5-11h-19V.5h19v12Z"/>`),
		g.Group(children),
	)
}