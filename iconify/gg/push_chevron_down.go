package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 7.414L6.414 6l5.657 5.657L17.728 6l1.414 1.414l-7.07 7.071L5 7.415Zm14 8.929H5v2h14v-2Z"/>`),
		g.Group(children),
	)
}