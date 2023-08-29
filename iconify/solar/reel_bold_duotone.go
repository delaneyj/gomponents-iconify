package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReelBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M2 12c0 5.523 4.477 10 10 10h9.25a.75.75 0 0 0 0-1.5h-3.98A9.993 9.993 0 0 0 22 12c0-5.523-4.477-10-10-10S2 6.477 2 12Z" clip-rule="evenodd" opacity=".5"/><path d="M16.5 10.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3ZM9 12a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Zm4.5-4.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 9a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/></g>`),
		g.Group(children),
	)
}