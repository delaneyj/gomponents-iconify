package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Orientation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 2.1c2 0 3 1.3 3 2.9h-1l1.5 2L16 5h-1c0-2.2-2-3.9-4-3.9V0L9 1.5L11 3v-.9z"/><path fill="currentColor" d="M9 9h6v6H8V0H0v16h16V8H9v1zM7 8H6v1h1v6H1V1h6v7z"/><path fill="currentColor" d="M2 8h1v1H2V8zm2 0h1v1H4V8z"/>`),
		g.Group(children),
	)
}