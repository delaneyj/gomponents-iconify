package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackAddTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 6.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Zm-4-2a.5.5 0 0 0-1 0V6H3.5a.5.5 0 0 0 0 1H5v1.5a.5.5 0 0 0 1 0V7h1.5a.5.5 0 0 0 0-1H6V4.5ZM5.5 12a5.5 5.5 0 0 0 4.9-8H12a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-1.257A5.477 5.477 0 0 0 5.5 12Zm-1.232 3A2 2 0 0 0 6 16h6a4 4 0 0 0 4-4V8a2 2 0 0 0-1-1.732V12a3 3 0 0 1-3 3H4.268Zm2 2A2 2 0 0 0 8 18h4a6 6 0 0 0 6-6v-2a2 2 0 0 0-1-1.732V12a5 5 0 0 1-5 5H6.268Z"/>`),
		g.Group(children),
	)
}