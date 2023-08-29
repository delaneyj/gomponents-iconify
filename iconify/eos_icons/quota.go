package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quota(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12V3h-1a10 10 0 1 0 10 10v-1Z"/><path fill="currentColor" d="M14 10V1a9 9 0 0 1 9 9Z"/>`),
		g.Group(children),
	)
}