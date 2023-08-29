package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDisplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 44H38C39.1046 44 40 43.1046 40 42V14H30V4H10C8.89543 4 8 4.89543 8 6V42C8 43.1046 8.89543 44 10 44Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 4L40 14"/><path fill="#43CCF8" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M24 34C28.9706 34 33 28 33 28C33 28 28.9706 22 24 22C19.0294 22 15 28 15 28C15 28 19.0294 34 24 34Z"/><path fill="#fff" d="M24 30C25.1046 30 26 29.1046 26 28C26 26.8954 25.1046 26 24 26C22.8954 26 22 26.8954 22 28C22 29.1046 22.8954 30 24 30Z"/></g>`),
		g.Group(children),
	)
}