package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openweather(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.672 30.636a17.694 17.694 0 1 0-33.344 0M4.5 30.635h39M8.791 35.817h15.706m4.393 0h2.982m4.392 0h2.982m-24.64 5.181h19.757"/>`),
		g.Group(children),
	)
}