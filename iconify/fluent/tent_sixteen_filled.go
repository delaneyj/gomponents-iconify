package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TentSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.405 1.353a.5.5 0 0 0-.768 0c-1.041 1.25-3.102 2.738-4.288 3.4a.5.5 0 0 0-.251.368L2.15 12h-.664a.5.5 0 0 0 0 1H14.5a.5.5 0 0 0 0-1h-.607l-.948-6.879a.5.5 0 0 0-.252-.368c-1.185-.662-3.246-2.15-4.288-3.4ZM5.944 12c.696-1.027 1.41-2.338 2.059-4.342c.59 1.815 1.246 3.108 2.06 4.342h-4.12Z"/>`),
		g.Group(children),
	)
}