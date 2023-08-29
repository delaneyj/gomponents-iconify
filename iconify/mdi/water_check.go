package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.75 16.25L17 21l-2.75-3l1.16-1.16L17 18.43l3.59-3.59l1.16 1.41M17.62 12C16.31 8.1 12 3.25 12 3.25S6 10 6 14c0 3.31 2.69 6 6 6h.34c-.22-.64-.34-1.3-.34-2c0-3.18 2.5-5.78 5.62-6Z"/>`),
		g.Group(children),
	)
}