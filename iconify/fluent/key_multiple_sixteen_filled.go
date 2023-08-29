package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyMultipleSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.96 1.03a3.5 3.5 0 0 0-3.886 4.189L1.146 8.146A.5.5 0 0 0 1 8.5v3a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-1h1.5A.5.5 0 0 0 6 10V8.707l.877-.877A4.977 4.977 0 0 1 6 5a4.99 4.99 0 0 1 1.96-3.97ZM7 5a4 4 0 1 1 6 3.465v1.828l.854.853a.5.5 0 0 1 0 .708l-.647.646l.647.646a.5.5 0 0 1-.054.754l-2 1.5a.5.5 0 0 1-.6 0l-2-1.5a.5.5 0 0 1-.2-.4V8.465A3.998 3.998 0 0 1 7 5Zm4.75-1a.75.75 0 1 0-1.5 0a.75.75 0 0 0 1.5 0Z"/>`),
		g.Group(children),
	)
}