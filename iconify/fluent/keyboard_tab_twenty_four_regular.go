package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardTabTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m18.03 11.47l-6.5-6.5a.75.75 0 0 0-1.133.976l.073.084l5.22 5.22H3.75a.75.75 0 0 0-.743.648L3 12c0 .38.282.694.648.743l.102.007h11.94l-5.22 5.22a.75.75 0 0 0-.073.976l.073.084a.75.75 0 0 0 .976.073l.084-.073l6.5-6.5a.75.75 0 0 0 .073-.976l-.073-.084l-6.5-6.5l6.5 6.5ZM21 18.5v-13a.75.75 0 0 0-1.5 0v13a.75.75 0 0 0 1.5 0Z"/>`),
		g.Group(children),
	)
}