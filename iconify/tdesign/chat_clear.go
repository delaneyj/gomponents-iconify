package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatClear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.5 2h21v16H6.876L1.5 22.704V2Zm2 2v14.296L6.124 16H20.5V4h-17Zm5.671 1.757L12 8.586l2.828-2.829l1.414 1.415L13.414 10l2.828 2.828l-1.414 1.415l-2.829-2.829l-2.828 2.829l-1.414-1.415L10.585 10L7.757 7.172L9.17 5.757Z"/>`),
		g.Group(children),
	)
}