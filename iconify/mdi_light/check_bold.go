package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 19l-5.657-5.657l2.121-2.121L9 14.757l8.485-8.485l2.122 2.121L9 19Zm-3.536-6.364l-.707.707L9 17.586l9.192-9.193l-.707-.707L9 16.172l-3.536-3.536Z"/>`),
		g.Group(children),
	)
}