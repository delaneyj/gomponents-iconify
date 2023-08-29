package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jionews(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m41.705 22.622l-12.846-7.416a1.591 1.591 0 0 0-2.387 1.378v14.832a1.591 1.591 0 0 0 2.387 1.378l12.846-7.416a1.591 1.591 0 0 0 0-2.756Zm-20.419 6.256H5.5m15.786-9.756H5.5m15.786-9.755H5.5m15.786 29.266H5.5"/>`),
		g.Group(children),
	)
}