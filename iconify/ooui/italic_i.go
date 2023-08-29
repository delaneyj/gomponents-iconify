package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ItalicI(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m5 19l.33-1.51l2.17-.66l2.9-13.66l-1.9-.63L9 1h7l-.71 1.6l-2.29.57l-2.83 13.66l2.14.66L12 19H5Z"/>`),
		g.Group(children),
	)
}