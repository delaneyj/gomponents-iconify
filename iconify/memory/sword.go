package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M8 3v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v-1h2v2h-1v1h-1v1h1v1h1v1h1v2h-1v1h-2v-1h-1v-1h-1v-1h-1v1h-1v1h-2v-2h1v-1h-1v-1h-1v-1H9v-1H8v-1H7v-1H6v-1H5V9H4V8H3V7H2V2h5v1h1M7 5H6V4H4v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h1v1h2v-1h-1v-1h-1v-1h-1V9h-1V8H9V7H8V6H7V5Z"/>`),
		g.Group(children),
	)
}