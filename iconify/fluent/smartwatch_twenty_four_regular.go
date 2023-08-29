package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartwatchTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 6.401V4a2 2 0 0 1 2-2h5a2 2 0 0 1 2 2v2.401A2.999 2.999 0 0 1 18 9v1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1v2a3 3 0 0 1-1.5 2.599V20a2 2 0 0 1-2 2h-5a2 2 0 0 1-2-2v-2.401A2.999 2.999 0 0 1 6 15V9a3 3 0 0 1 1.5-2.599Zm2-2.901A.5.5 0 0 0 9 4v2h6V4a.5.5 0 0 0-.5-.5h-5Zm7 11.5V9A1.5 1.5 0 0 0 15 7.5H9A1.5 1.5 0 0 0 7.5 9v6A1.5 1.5 0 0 0 9 16.5h6a1.5 1.5 0 0 0 1.5-1.5ZM9 20a.5.5 0 0 0 .5.5h5a.5.5 0 0 0 .5-.5v-2H9v2Z"/>`),
		g.Group(children),
	)
}