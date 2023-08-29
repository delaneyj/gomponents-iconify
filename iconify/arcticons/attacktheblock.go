package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attacktheblock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.38 5.5H24v24.13H5.98L22.38 5.5zM5.98 29.63V42.5L17.5 29.63m8.88 0l3.84 4.3H17.78l3.85-4.3M25.62 5.5H24v24.13h18.02L25.62 5.5zm16.4 24.13V42.5L30.5 29.63"/>`),
		g.Group(children),
	)
}