package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Landsbanken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.405 13.72l-3.071-5.318H12.557l3.06 5.301Zm1.295 2.025l14.8 23.821H26.194l2.943-5.098H15.048l-2.961 5.13H4.5L14.617 21.07l-3.316-5.348Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.572 25.572l2.075 3.557l-4.118.017l-4.12.018l2.046-3.574l2.043-3.576Z"/>`),
		g.Group(children),
	)
}