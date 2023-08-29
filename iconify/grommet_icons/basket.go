package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Basket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M2 12h20l-2 11H4L2 12Zm18-4l-6-7M4 8l6-7M1 8h22v4H1V8Zm7 7v5m8-5v5m-4-5v5"/>`),
		g.Group(children),
	)
}