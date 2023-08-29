package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M11.092 2.581a1 1 0 0 1 1.754-.116l.062.116l8.005 17.365c.198.566.05 1.196-.378 1.615a1.53 1.53 0 0 1-1.459.393l-7.077-2.398L5.1 21.894a1.535 1.535 0 0 1-1.52-.231l-.112-.1c-.398-.386-.556-.954-.393-1.556l.047-.15l7.97-17.276z"/></g>`),
		g.Group(children),
	)
}