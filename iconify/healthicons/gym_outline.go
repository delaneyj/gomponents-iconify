package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GymOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M19 32V16h-3v16h3Zm-5-3V19h-3v4h-1v2h1v4h3Zm15-13v16h3V16h-3Zm5 13h3v-4h1v-2h-1v-4h-3v10Zm-13-4h6v-2h-6v2Z"/><path fill-rule="evenodd" d="M6 9a3 3 0 0 1 3-3h30a3 3 0 0 1 3 3v30a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3V9Zm2 0a1 1 0 0 1 1-1h30a1 1 0 0 1 1 1v30a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1V9Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}