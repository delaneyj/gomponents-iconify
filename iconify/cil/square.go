package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Square(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M476 16H36a20.023 20.023 0 0 0-20 20v440a20.023 20.023 0 0 0 20 20h440a20.023 20.023 0 0 0 20-20V36a20.023 20.023 0 0 0-20-20Zm-12 448H48V48h416Z"/>`),
		g.Group(children),
	)
}