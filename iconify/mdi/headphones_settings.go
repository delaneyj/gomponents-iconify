package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphonesSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1a9 9 0 0 1 9 9v7a3 3 0 0 1-3 3h-3v-8h4v-2a7 7 0 0 0-7-7a7 7 0 0 0-7 7v2h4v8H6a3 3 0 0 1-3-3v-7a9 9 0 0 1 9-9m3 23v-2h2v2h-2m-4 0v-2h2v2h-2m-4 0v-2h2v2H7Z"/>`),
		g.Group(children),
	)
}