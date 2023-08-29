package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BombMinimalisticOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M15.665 2.83a.75.75 0 0 1 1.006.335l.5 1a.75.75 0 1 1-1.342.67l-.5-1a.75.75 0 0 1 .336-1.006Z"/><path fill-rule="evenodd" d="M1.25 14.5a8.25 8.25 0 0 1 13.53-6.34l1.69-1.69a.75.75 0 1 1 1.06 1.06l-1.69 1.69A8.25 8.25 0 1 1 1.25 14.5ZM9.5 7.75a6.75 6.75 0 1 0 0 13.5a6.75 6.75 0 0 0 0-13.5Z" clip-rule="evenodd"/><path d="M19.835 6.83a.75.75 0 1 0-.67 1.341l1 .5a.75.75 0 1 0 .67-1.342l-1-.5Zm.695-2.3a.75.75 0 0 0-1.06-1.06l-1 1a.75.75 0 0 0 1.06 1.06l1-1Z"/></g>`),
		g.Group(children),
	)
}