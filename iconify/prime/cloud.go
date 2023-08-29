package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 18.75h-7a6.75 6.75 0 1 1 6.2-9.42a4.43 4.43 0 0 1 .8-.08A5.07 5.07 0 0 1 21.25 14a4.75 4.75 0 0 1-4.75 4.75Zm-7-12a5.25 5.25 0 0 0 0 10.5h7A3.26 3.26 0 0 0 19.75 14a3.57 3.57 0 0 0-3.25-3.25a3.09 3.09 0 0 0-1 .19a.78.78 0 0 1-.58 0a.73.73 0 0 1-.37-.44A5.24 5.24 0 0 0 9.5 6.75Z"/>`),
		g.Group(children),
	)
}