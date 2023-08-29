package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.5 2h21v16H6.876L1.5 22.704V2Zm2 2v14.296L6.124 16H20.5V4h-17Zm14.157 3.586l-7.071 7.07l-4.243-4.242L7.757 9l2.829 2.828l5.657-5.657l1.414 1.415Z"/>`),
		g.Group(children),
	)
}