package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M36 416h440a20.023 20.023 0 0 0 20-20V116a20.023 20.023 0 0 0-20-20H36a20.023 20.023 0 0 0-20 20v280a20.023 20.023 0 0 0 20 20Zm12-288h416v256H48Z"/>`),
		g.Group(children),
	)
}