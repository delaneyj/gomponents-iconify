package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ErrorSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M9.177 10.105L8 8.928l-1.177 1.177l-.928-.928L7.072 8L5.895 6.823l.928-.928L8 7.072l1.177-1.177l.928.928L8.928 8l1.177 1.177l-.928.928Z"/><path d="M12 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm-1 0a3 3 0 1 0-6 0a3 3 0 0 0 6 0Z"/></g>`),
		g.Group(children),
	)
}