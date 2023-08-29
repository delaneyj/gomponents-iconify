package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumbersFourOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.989 4H15.5v10h2v2h-2v4h-2v-4H6v-2.323L12.989 4Zm.511 10V6.708L8.234 14H13.5Z"/>`),
		g.Group(children),
	)
}