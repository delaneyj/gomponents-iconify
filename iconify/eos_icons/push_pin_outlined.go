package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushPinOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.996 12.011A3.005 3.005 0 0 1 16 9V4h1l-.013.011A1.023 1.023 0 0 0 18 3a1.007 1.007 0 0 0-1.015-.988L17 2H7a1 1 0 0 0-.003 2H8l-.004 5A2.991 2.991 0 0 1 5 12v2h6v7l1 1l1-1v-7h6v-2ZM9 12a4.941 4.941 0 0 0 1-3V4h4v5a4.99 4.99 0 0 0 1 3Z"/>`),
		g.Group(children),
	)
}