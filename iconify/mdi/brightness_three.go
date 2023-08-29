package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 2c-1.05 0-2.05.16-3 .46c4.06 1.27 7 5.04 7 9.54c0 4.5-2.94 8.27-7 9.54c.95.3 1.95.46 3 .46a10 10 0 0 0 10-10A10 10 0 0 0 9 2Z"/>`),
		g.Group(children),
	)
}