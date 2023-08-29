package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlaskTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.1 15h11.8M14 3v7.342A6 6 0 0 1 15.318 21H8.683A6 6 0 0 1 10 10.34V3h4zM9 3h6"/>`),
		g.Group(children),
	)
}