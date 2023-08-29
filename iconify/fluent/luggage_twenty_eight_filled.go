package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuggageTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 2.25a.75.75 0 0 1 .75-.75h8.5a.75.75 0 0 1 0 1.5H18v1h1.5A3.5 3.5 0 0 1 23 7.5V21a3.501 3.501 0 0 1-3 3.465v.785a.75.75 0 0 1-1.5 0v-.75h-9v.75a.75.75 0 0 1-1.5 0v-.785A3.5 3.5 0 0 1 5 21V7.5A3.5 3.5 0 0 1 8.5 4H10V3h-.25A.75.75 0 0 1 9 2.25Zm2.5.75v1h5V3h-5ZM9.75 9a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5h-8.5Z"/>`),
		g.Group(children),
	)
}