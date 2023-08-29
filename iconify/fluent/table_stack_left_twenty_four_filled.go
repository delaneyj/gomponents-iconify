package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableStackLeftTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 21h1.75A3.25 3.25 0 0 0 21 17.75V16h-5v5Zm0-6.5h5v-5h-5v5Zm-1.5-5v5h-5v-5h5ZM16 8h5V6.25A3.25 3.25 0 0 0 17.75 3H16v5Zm-1.5-5v5h-5V3.75a.75.75 0 0 1 .75-.75h4.25Zm0 13v5h-4.25a.75.75 0 0 1-.75-.75V16h5ZM4.5 3.75a.75.75 0 0 0-1.5 0v16.5a.75.75 0 0 0 1.5 0V3.75Z"/>`),
		g.Group(children),
	)
}