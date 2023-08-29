package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 17.5A4.505 4.505 0 0 1 8.5 13a1 1 0 1 0-2 0c0 3.584 2.916 6.5 6.5 6.5s6.5-2.916 6.5-6.5s-2.916-6.5-6.5-6.5a1 1 0 1 0 0 2c2.481 0 4.5 2.019 4.5 4.5s-2.019 4.5-4.5 4.5zM10.656 4a1 1 0 0 1 0 2H7.413l1.708 1.707l4.093 4.092a1.001 1.001 0 0 1-1.414 1.416L7.707 9.121L6 7.413v3.243a1 1 0 0 1-2 0V4h6.656"/>`),
		g.Group(children),
	)
}