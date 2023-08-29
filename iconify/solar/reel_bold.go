package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReelBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12c0 5.523 4.477 10 10 10h9.25a.75.75 0 0 0 0-1.5h-3.98A9.993 9.993 0 0 0 22 12c0-5.523-4.477-10-10-10S2 6.477 2 12Zm10-3a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 9a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm-4.5-7.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3ZM18 12a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}