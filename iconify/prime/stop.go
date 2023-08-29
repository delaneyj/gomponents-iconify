package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 19.75H7A2.75 2.75 0 0 1 4.25 17V7A2.75 2.75 0 0 1 7 4.25h10A2.75 2.75 0 0 1 19.75 7v10A2.75 2.75 0 0 1 17 19.75Zm-10-14A1.25 1.25 0 0 0 5.75 7v10A1.25 1.25 0 0 0 7 18.25h10A1.25 1.25 0 0 0 18.25 17V7A1.25 1.25 0 0 0 17 5.75Z"/>`),
		g.Group(children),
	)
}