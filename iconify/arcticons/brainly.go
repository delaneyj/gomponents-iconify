package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brainly(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-miterlimit="5.5" d="m9.332 8.28l22.827-4.172M15.896 43.841l22.827-4.173"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.617 23.054l-9.242 1.696m9.242-1.696a4.482 4.482 0 0 1 5.601 4.403a6.956 6.956 0 0 1-5.601 6.458l-9.242 1.696V13.889l9.242-1.696a4.482 4.482 0 0 1 5.601 4.403a6.956 6.956 0 0 1-5.601 6.458Zm-10.836-8.873l1.594-.292m-1.594 22.014l1.594-.293"/>`),
		g.Group(children),
	)
}