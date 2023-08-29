package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileUserFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 2l5 5v14.008a.993.993 0 0 1-.993.992H3.993A1 1 0 0 1 3 21.008V2.992C3 2.444 3.445 2 3.993 2H16Zm-4 9.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM7.527 17h8.945a4.5 4.5 0 0 0-8.945 0Z"/>`),
		g.Group(children),
	)
}