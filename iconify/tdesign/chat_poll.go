package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatPoll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.5 2h21v16H6.876L1.5 22.704V2Zm2 2v14.296L6.124 16H20.5V4h-17ZM13 6v8h-2V6h2Zm5 3v5h-2V9h2ZM8 11v3H6v-3h2Z"/>`),
		g.Group(children),
	)
}