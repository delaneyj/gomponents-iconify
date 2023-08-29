package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picsart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5A14.1 14.1 0 0 0 9.9 18.6a14.1 14.1 0 0 0 24.07 9.97a14.1 14.1 0 0 0 0-19.94A14.27 14.27 0 0 0 24 4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.932 12.29a6.28 6.28 0 1 0 4.44 1.84a6.356 6.356 0 0 0-4.44-1.84Zm-6.28 18.884v8.45a3.876 3.876 0 0 1-3.876 3.876h0A3.876 3.876 0 0 1 9.9 39.624V18.6"/>`),
		g.Group(children),
	)
}