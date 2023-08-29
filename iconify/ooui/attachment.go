package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attachment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 19.75a4.25 4.25 0 0 1-4.25-4.25V9a2.75 2.75 0 0 1 5.5 0v6h-1.5V9a1.25 1.25 0 0 0-2.5 0v6.5a2.75 2.75 0 0 0 5.5 0V4a2.25 2.25 0 0 0-4.5 0v1h-1.5V4a3.75 3.75 0 0 1 7.5 0v11.5a4.25 4.25 0 0 1-4.25 4.25z"/>`),
		g.Group(children),
	)
}