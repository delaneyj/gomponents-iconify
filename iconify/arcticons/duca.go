package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Duca(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.167 5.5v24.667H42.5V17.833L30.167 5.5zm0 0H5.5m24.667 3.083H5.5m24.667 3.084H5.5m24.667 3.083H5.5m24.667 3.083H5.5m24.667 12.334H5.5m33.917 3.083H5.5m30.833 3.083H5.5m27.75 3.084H5.5M30.167 42.5H5.5"/>`),
		g.Group(children),
	)
}