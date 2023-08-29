package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatMessage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.5 22.704V2h21v16H6.876L1.5 22.704Zm2-4.408L6.124 16H20.5V4h-17v14.296ZM13.004 11H11V8.996h2.004V11Zm-5 0H6V8.996h2.004V11Zm10 0H16V8.996h2.004V11Z"/>`),
		g.Group(children),
	)
}