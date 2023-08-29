package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClearFormatting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.031 1.888l9.657 9.657l-8.345 8.345l-.27.272H20v2H6.748L.253 15.667L14.03 1.888Zm.322 16.164l6.507-6.507l-6.829-6.828l-6.828 6.828l6.828 6.828l.32-.323l.002.002ZM5.788 12.96l-2.707 2.707l4.495 4.495h4.68l.365-.37l-6.833-6.833Z"/>`),
		g.Group(children),
	)
}