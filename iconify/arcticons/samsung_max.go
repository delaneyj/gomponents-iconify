package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SamsungMax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 25.04l1.85 15.955h10.174l1.233-10.675l2.736 10.675h9.25l2.35-6.898l1.657 6.898h9.75l-4.547-23.007H28.74l-2.35 7.322l-3.854-18.344H12.207l-3.66 23.277Zm4.046 5.203l7.978 10.79M12.207 6.966l5.55 23.354m8.633-5.01l3.352 15.608m-1.002-22.93l3.56 16.994"/>`),
		g.Group(children),
	)
}