package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoardTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.5 3v6.5H3V6.75A3.75 3.75 0 0 1 6.75 3h6.75ZM15 3v13.5h10V6.75A3.75 3.75 0 0 0 21.25 3H15Zm10 15H15v7h6.25A3.75 3.75 0 0 0 25 21.25V18Zm-11.5 7V11H3v10.25A3.75 3.75 0 0 0 6.75 25h6.75Z"/>`),
		g.Group(children),
	)
}