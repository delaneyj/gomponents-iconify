package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Search(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M18.425 18.425L23.5 23.5m-12.5-2C5.201 21.5.5 16.799.5 11S5.201.5 11 .5S21.5 5.201 21.5 11S16.799 21.5 11 21.5Z"/>`),
		g.Group(children),
	)
}