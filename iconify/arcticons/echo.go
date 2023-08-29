package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Echo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.84 2.75a152.53 152.53 0 0 0-1.79 17.88a69.39 69.39 0 0 1 7.46-.6h2.36l-.48 2.32c-.2 1-.29 1.27-.49 2.1a38.71 38.71 0 0 1 4.8.61l2.37.68l-1.2 2.15c-3.13 5.63-8.73 8-10.69 8.72A73 73 0 0 0 24.32 44c.12.57.21 1 .32 1.5M28.58 8a3.69 3.69 0 0 1 0 7.37h0a3.69 3.69 0 0 1-3.69-3.69h0A3.69 3.69 0 0 1 28.58 8Z"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}