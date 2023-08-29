package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmstripPlaySixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 3A2.5 2.5 0 0 0 1 5.5v5A2.5 2.5 0 0 0 3.5 13h9a2.5 2.5 0 0 0 2.5-2.5v-5A2.5 2.5 0 0 0 12.5 3h-9Zm9 2a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-1 0v-1a.5.5 0 0 1 .5-.5ZM12 9.5a.5.5 0 0 1 1 0v1a.5.5 0 0 1-1 0v-1ZM3.5 5a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-1 0v-1a.5.5 0 0 1 .5-.5ZM3 9.5a.5.5 0 0 1 1 0v1a.5.5 0 0 1-1 0v-1Zm4-3c0-.385.346-.626.62-.43l2.206 1.56c.233.165.233.573 0 .738L7.621 9.93C7.346 10.126 7 9.885 7 9.499V6.501Z"/>`),
		g.Group(children),
	)
}