package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuLeftOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 18h-1.5l-6-6l6-6H15v12m-4.67-6L13 14.67V9.33L10.33 12Z"/>`),
		g.Group(children),
	)
}