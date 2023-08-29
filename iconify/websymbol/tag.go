package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 136v270l-595 595L0 596L595 1h270zM840 296q28-28 28-67.5T840 161t-67.5-28t-67.5 28t-28 67.5t28 67.5t67.5 28t67.5-28z"/>`),
		g.Group(children),
	)
}