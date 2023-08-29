package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M.5 1a.5.5 0 0 0-.5.5v10a.5.5 0 0 0 .5.5H3v-1H1V2h13v9h-2v1h2.5a.5.5 0 0 0 .5-.5v-10a.5.5 0 0 0-.5-.5H.5Z"/><path fill="currentColor" d="M7.854 9.146a.5.5 0 0 0-.708 0l-4 4A.5.5 0 0 0 3.5 14h8a.5.5 0 0 0 .354-.854l-4-4Z"/>`),
		g.Group(children),
	)
}