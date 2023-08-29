package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckboxBlankOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.84 22.73L19.1 21H5a2 2 0 0 1-2-2V4.9L1.11 3l1.28-1.27l19.72 19.73l-1.27 1.27M21 5a2 2 0 0 0-2-2H6.2L21 17.8V5Z"/>`),
		g.Group(children),
	)
}