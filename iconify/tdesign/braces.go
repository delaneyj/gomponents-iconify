package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Braces(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.5 17.6A3.4 3.4 0 0 0 5.9 21H7v-2H5.9a1.4 1.4 0 0 1-1.4-1.4V14c0-.768-.288-1.47-.764-2c.476-.53.764-1.232.764-2V6.4A1.4 1.4 0 0 1 5.9 5H7V3H5.9a3.4 3.4 0 0 0-3.4 3.4V10a1 1 0 0 1-1 1H.4v2h1.1a1 1 0 0 1 1 1v3.6ZM17 21h1.1a3.4 3.4 0 0 0 3.4-3.4V14a1 1 0 0 1 1-1h1.1v-2h-1.1a1 1 0 0 1-1-1V6.4A3.4 3.4 0 0 0 18.1 3H17v2h1.1a1.4 1.4 0 0 1 1.4 1.4V10c0 .768.288 1.47.764 2a2.989 2.989 0 0 0-.764 2v3.6a1.4 1.4 0 0 1-1.4 1.4H17v2Z"/>`),
		g.Group(children),
	)
}