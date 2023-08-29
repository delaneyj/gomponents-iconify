package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadSide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.2 2c-2-.1-4 .7-5.5 2.1c-1.4 1.4-2.2 3.4-2.2 5.4l-1.9 4c-.2.5 0 1.1.5 1.3c.1.2.3.2.4.2h1v2c0 1.1.9 2 2 2h1v2c0 .6.4 1 1 1h9.3c.5-.2.8-.7.7-1.2l-.9-3.2l1.9-7.3v-.5c0-4.1-3.2-7.6-7.3-7.8z"/>`),
		g.Group(children),
	)
}