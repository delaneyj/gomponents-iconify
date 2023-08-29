package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beatfind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.188 26.192l-4.225 16.901a.315.315 0 0 0 .548.277l16.967-20.56a.63.63 0 0 0-.486-1.03H21.938c-3.51 0-4.74-1.755-3.437-4.061"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.648 21.78l4.225-16.901a.315.315 0 0 0-.548-.277L12.36 25.162a.63.63 0 0 0 .485 1.03H25.9c3.51 0 4.74 1.755 3.436 4.061"/>`),
		g.Group(children),
	)
}