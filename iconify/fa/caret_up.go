package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1024 960q0 26-19 45t-45 19H64q-26 0-45-19T0 960t19-45l448-448q19-19 45-19t45 19l448 448q19 19 19 45z"/>`),
		g.Group(children),
	)
}