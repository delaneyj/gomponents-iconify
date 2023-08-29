package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReelTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 2.75a9.25 9.25 0 1 0 0 18.5a9.25 9.25 0 0 0 0-18.5Zm5.48 18.5A10.745 10.745 0 0 0 22.75 12c0-5.937-4.813-10.75-10.75-10.75S1.25 6.063 1.25 12S6.063 22.75 12 22.75h10a.75.75 0 0 0 0-1.5h-4.52ZM12 9.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5ZM8.25 12a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0Z" clip-rule="evenodd"/><path d="M19.5 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-13 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM12 4.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm0 13a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z"/></g>`),
		g.Group(children),
	)
}