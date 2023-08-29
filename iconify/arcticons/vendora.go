package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vendora(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m27.247 40.645l14.745-25.54a3.749 3.749 0 0 0-3.246-5.624H9.254a3.749 3.749 0 0 0-3.246 5.624l14.745 25.54a3.749 3.749 0 0 0 6.494 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.254 9.481c14.746 0 26.502 16.425 19.392 28.74"/>`),
		g.Group(children),
	)
}