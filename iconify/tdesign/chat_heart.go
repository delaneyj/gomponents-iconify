package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.5 2h21v16H6.876L1.5 22.704V2Zm2 2v14.296L6.124 16H20.5V4h-17ZM12 5.772a3.251 3.251 0 0 0-4.065 5.026L12 14.864l4.065-4.066A3.25 3.25 0 0 0 12 5.772Zm-.883 1.844L12 8.5l.884-.884a1.25 1.25 0 0 1 1.768 1.768l-2.651 2.652l-2.652-2.652a1.25 1.25 0 0 1 1.768-1.768Z"/>`),
		g.Group(children),
	)
}