package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadWeb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M28 40H7C5.34315 40 4 38.6569 4 37V11C4 9.34315 5.34315 8 7 8H41C42.6569 8 44 9.34315 44 11V23.0588"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M39 41L39 29"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M34 34L39 29L44 34"/><path fill="#2F88FF" stroke="#000" stroke-width="4" d="M4 11C4 9.34315 5.34315 8 7 8H41C42.6569 8 44 9.34315 44 11V20H4V11Z"/><circle r="2" fill="#fff" transform="matrix(0 -1 -1 0 10 14)"/><circle r="2" fill="#fff" transform="matrix(0 -1 -1 0 16 14)"/></g>`),
		g.Group(children),
	)
}