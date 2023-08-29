package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwitchSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M.5 0a.5.5 0 0 0-.5.5v11a.5.5 0 0 0 .5.5h2v2.5a.5.5 0 0 0 .825.38L6.685 12H11.5a.5.5 0 0 0 .354-.146l3-3A.5.5 0 0 0 15 8.5v-8a.5.5 0 0 0-.5-.5H.5ZM10 8V3h1v5h-1ZM7 3v5h1V3H7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}