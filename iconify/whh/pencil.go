package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m780 949l-172-53l288-288l53 172q-61 18-106 63t-63 106zM32 352Q0 319 0 271.5T33 192L192 33q32-33 79.5-33T352 32l512 512l-320 320zm937 493l55 179l-179-55q16-43 48.5-75.5T969 845z"/>`),
		g.Group(children),
	)
}