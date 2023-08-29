package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeDownTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 6a1 1 0 0 1 1 1v11.585l1.293-1.292a1 1 0 0 1 1.32-.083l.094.083a1 1 0 0 1 .083 1.32l-.083.094l-3 3a1 1 0 0 1-1.32.083l-.094-.083l-3-3a1 1 0 0 1 1.32-1.497l.094.083L11 18.585V7a1 1 0 0 1 1-1Zm0-4a5 5 0 0 1 2 9.584V9.872a3.5 3.5 0 1 0-4 0v1.712A5.001 5.001 0 0 1 12 2Z"/>`),
		g.Group(children),
	)
}