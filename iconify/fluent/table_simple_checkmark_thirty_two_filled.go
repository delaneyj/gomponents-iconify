package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleCheckmarkThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 3H7.5A4.5 4.5 0 0 0 3 7.5V15h12V3ZM3 24.5V17h12v12H7.5A4.5 4.5 0 0 1 3 24.5ZM17 29V17h12v7.5a4.5 4.5 0 0 1-4.5 4.5H17ZM29 7.5V15H17V3h7.5A4.5 4.5 0 0 1 29 7.5Zm-2.543 14.457a1 1 0 0 0-1.414-1.414L22 23.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0l3.75-3.75Z"/>`),
		g.Group(children),
	)
}