package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppsTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 11.5a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm0 10a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm10-10a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm0 10a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm-10-12a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 10a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm10-10a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 10a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/>`),
		g.Group(children),
	)
}