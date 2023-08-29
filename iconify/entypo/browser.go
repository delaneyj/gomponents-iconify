package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Browser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 2H2C.9 2 0 2.9 0 4v12c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zM4.5 3.75a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5zm-2.75.75a.75.75 0 1 1 1.5 0a.75.75 0 0 1-1.5 0zM18 16H2V7h16v9zm0-11H6V4h12.019L18 5z"/>`),
		g.Group(children),
	)
}