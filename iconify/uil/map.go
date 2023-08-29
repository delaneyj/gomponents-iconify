package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.32 5.05l-6-2h-.07a.7.7 0 0 0-.14 0h-.43L9 5L3.32 3.05a1 1 0 0 0-.9.14A1 1 0 0 0 2 4v14a1 1 0 0 0 .68.95l6 2a1 1 0 0 0 .62 0l5.7-1.9L20.68 21a1.19 1.19 0 0 0 .32 0a.94.94 0 0 0 .58-.19A1 1 0 0 0 22 20V6a1 1 0 0 0-.68-.95ZM8 18.61l-4-1.33V5.39l4 1.33Zm6-1.33l-4 1.33V6.72l4-1.33Zm6 1.33l-4-1.33V5.39l4 1.33Z"/>`),
		g.Group(children),
	)
}