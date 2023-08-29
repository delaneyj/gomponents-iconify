package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbstractOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.732 11l-4.003-7a1.967 1.967 0 0 0-1.72-1H8.005a2.038 2.038 0 0 0-1.733 1l-4.003 7a1.999 1.999 0 0 0 0 2l4.003 7a2.038 2.038 0 0 0 1.733 1h8.006a1.966 1.966 0 0 0 1.719-1l4.003-7a1.999 1.999 0 0 0 0-2ZM16 19H8l-4-7l4-7h8l4 7Z"/>`),
		g.Group(children),
	)
}