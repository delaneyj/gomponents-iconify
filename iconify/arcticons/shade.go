package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.98 5.945L4.5 24.968h5.571v17.087h9.296V30.232h9.266v11.823h9.296V24.968H43.5L23.98 5.945z"/>`),
		g.Group(children),
	)
}