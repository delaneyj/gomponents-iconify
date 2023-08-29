package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Osano(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6.091A5.909 5.909 0 1 0 17.909 12A5.91 5.91 0 0 0 12 6.091M12 0A12 12 0 1 1 0 12A12 12 0 0 1 12 0z"/>`),
		g.Group(children),
	)
}