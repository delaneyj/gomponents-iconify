package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleSwitchOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 8h4v1h1v4H9v1H5v-1H4V9h1V8m14-3v1h1v1h1v8h-1v1h-1v1H3v-1H2v-1H1V7h1V6h1V5h16m-1 2H4v1H3v6h1v1h14v-1h1V8h-1V7Z"/>`),
		g.Group(children),
	)
}