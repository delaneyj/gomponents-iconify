package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureInPictureBottomRightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 11h-8v6h8v-6m-2 4h-4v-2h4v2m4-12H3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V5c0-1.12-.9-2-2-2m0 16H3V4.97h18V19Z"/>`),
		g.Group(children),
	)
}