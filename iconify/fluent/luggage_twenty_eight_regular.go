package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LuggageTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 9.75A.75.75 0 0 1 9.75 9h8.5a.75.75 0 0 1 0 1.5h-8.5A.75.75 0 0 1 9 9.75Zm.75-8.25a.75.75 0 0 0 0 1.5H10v1H8.5A3.5 3.5 0 0 0 5 7.5V21a3.5 3.5 0 0 0 3 3.465v.785a.75.75 0 0 0 1.5 0v-.75h9v.75a.75.75 0 0 0 1.5 0v-.785A3.501 3.501 0 0 0 23 21V7.5A3.5 3.5 0 0 0 19.5 4H18V3h.25a.75.75 0 0 0 0-1.5h-8.5ZM16.5 3v1h-5V3h5Zm-8 2.5h11a2 2 0 0 1 2 2V21a2 2 0 0 1-2 2h-11a2 2 0 0 1-2-2V7.5a2 2 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}