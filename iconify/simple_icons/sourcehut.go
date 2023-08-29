package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sourcehut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C5.371 0 0 5.371 0 12s5.371 12 12 12s12-5.371 12-12S18.629 0 12 0Zm0 21.677A9.675 9.675 0 0 1 2.323 12A9.675 9.675 0 0 1 12 2.323A9.675 9.675 0 0 1 21.677 12A9.675 9.675 0 0 1 12 21.677Z"/>`),
		g.Group(children),
	)
}