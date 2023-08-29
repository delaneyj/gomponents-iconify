package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dumbbell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4.5v7h1v-7H4ZM3.5 3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V8.75h3V12a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v3.25h-3V4a1 1 0 0 0-1-1h-2ZM11 4.5v7h1v-7h-1Zm4.25.75A.75.75 0 0 1 16 6v4.5a.75.75 0 0 1-1.5 0V6a.75.75 0 0 1 .75-.75ZM1.5 6A.75.75 0 0 0 0 6v4.5a.75.75 0 0 0 1.5 0V6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}