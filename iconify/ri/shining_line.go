package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShiningLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 11c5.523 0 10-4.477 10-10h2c0 5.523 4.477 10 10 10v2c-5.523 0-10 4.477-10 10h-2c0-5.523-4.477-10-10-10v-2Zm4.803 1A12.044 12.044 0 0 1 12 18.197A12.043 12.043 0 0 1 18.197 12A12.044 12.044 0 0 1 12 5.803A12.044 12.044 0 0 1 5.803 12Z"/>`),
		g.Group(children),
	)
}