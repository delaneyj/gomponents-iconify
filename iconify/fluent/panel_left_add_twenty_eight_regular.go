package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanelLeftAddTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.75 4A3.75 3.75 0 0 0 2 7.75v11.5A3.75 3.75 0 0 0 5.75 23h7.673a7.451 7.451 0 0 1-.36-1.5h-2.567v-16h11.75a2.25 2.25 0 0 1 2.25 2.25v6.405c.554.35 1.058.769 1.5 1.247V7.75A3.75 3.75 0 0 0 22.247 4H5.75ZM3.5 7.75A2.25 2.25 0 0 1 5.75 5.5h3.246v16H5.75a2.25 2.25 0 0 1-2.25-2.25V7.75ZM26.996 20.5a6.5 6.5 0 1 1-13 0a6.5 6.5 0 0 1 13 0Zm-6-4a.5.5 0 0 0-1 0V20h-3.5a.5.5 0 0 0 0 1h3.5v3.5a.5.5 0 0 0 1 0V21h3.5a.5.5 0 0 0 0-1h-3.5v-3.5Z"/>`),
		g.Group(children),
	)
}