package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleSwitchOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 5h16v1h1v1h1v8h-1v1h-1v1H3v-1H2v-1H1V7h1V6h1V5m10 3v1h-1v4h1v1h4v-1h1V9h-1V8h-4Z"/>`),
		g.Group(children),
	)
}