package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TypeDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M4 12C4 10.8954 4.89543 10 6 10H42C43.1046 10 44 10.8954 44 12V36C44 37.1046 43.1046 38 42 38H6C4.89543 38 4 37.1046 4 36V12Z"/><path fill="#2F88FF" stroke="#fff" d="M15.4142 32.5858C15.7893 32.2107 16.298 32 16.8284 32H30.2602C30.7376 32 31.1992 32.1708 31.5617 32.4815L38 38H10L15.4142 32.5858Z"/><path stroke="#000" d="M39 38H9"/><circle cx="15" cy="21" r="4" fill="#43CCF8" stroke="#fff"/><circle cx="33" cy="21" r="4" fill="#43CCF8" stroke="#fff"/><path stroke="#fff" d="M15 25L33 25"/><path stroke="#fff" d="M15 17L33 17"/></g>`),
		g.Group(children),
	)
}