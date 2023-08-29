package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M15 4.5H3a.5.5 0 0 0-.5.5c0 4.668 2.874 8.5 6.5 8.5s6.5-3.832 6.5-8.5a.5.5 0 0 0-.5-.5Zm-.512 1c-.19 3.932-2.608 7-5.488 7c-2.88 0-5.298-3.068-5.488-7h10.976Z" clip-rule="evenodd"/><path d="m13.862 10.329l.276-.961c.303.086.63.132.97.132C16.46 9.5 17.5 8.78 17.5 8s-1.04-1.5-2.393-1.5v-1C16.95 5.5 18.5 6.572 18.5 8s-1.55 2.5-3.393 2.5c-.431 0-.852-.058-1.245-.171ZM3.5 14.75h11a.75.75 0 0 1 0 1.5h-11a.75.75 0 0 1 0-1.5Z"/></g>`),
		g.Group(children),
	)
}