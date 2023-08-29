package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 14.5v-13h9v13H11V11a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v3.5H3.5Zm3 0h3v-3h-3v3ZM2 1a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V1Zm3 2.75A.75.75 0 0 1 5.75 3h1a.75.75 0 0 1 0 1.5h-1A.75.75 0 0 1 5 3.75ZM9.25 3a.75.75 0 0 0 0 1.5h1a.75.75 0 0 0 0-1.5h-1ZM5 7.25a.75.75 0 0 1 .75-.75h1a.75.75 0 0 1 0 1.5h-1A.75.75 0 0 1 5 7.25Zm4.25-.75a.75.75 0 0 0 0 1.5h1a.75.75 0 0 0 0-1.5h-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}