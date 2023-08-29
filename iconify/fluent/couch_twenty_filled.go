package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CouchTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 4H6a2 2 0 0 0-2 2v1a3 3 0 0 1 2.829 2h6.342A3 3 0 0 1 16 7V6a2 2 0 0 0-2-2Zm2 4a2 2 0 0 0-1.938 1.505c-.068.268-.286.495-.562.495h-7c-.276 0-.494-.227-.562-.495A2 2 0 0 0 2 10v2a2 2 0 0 0 2 2v1.5a.5.5 0 0 0 1 0V14h10v1.5a.5.5 0 0 0 1 0V14a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}