package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormMultipleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path d="M7.25 11.5a.75.75 0 1 0 0 1.5a.75.75 0 0 0 0-1.5Z"/><path d="M6.25 3A3.25 3.25 0 0 0 3 6.25v9a3.25 3.25 0 0 0 3.25 3.25h9a3.25 3.25 0 0 0 3.25-3.25v-9A3.25 3.25 0 0 0 15.25 3h-9ZM5 12.25a2.25 2.25 0 1 1 4.5 0a2.25 2.25 0 0 1-4.5 0Zm5.5 0a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1-.75-.75ZM5 7.75A.75.75 0 0 1 5.75 7h10a.75.75 0 0 1 0 1.5h-10A.75.75 0 0 1 5 7.75Z"/><path d="M8.75 21a3.247 3.247 0 0 1-2.74-1.5h9.74a3.75 3.75 0 0 0 3.75-3.75V6.011a3.248 3.248 0 0 1 1.5 2.74v7A5.25 5.25 0 0 1 15.75 21h-7Z"/></g>`),
		g.Group(children),
	)
}