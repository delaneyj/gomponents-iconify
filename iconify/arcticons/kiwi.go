package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kiwi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.85 15.47s-4.957-8.336-17.726-8.31C11.354 7.187 3.85 15.128 3.85 24s7.641 16.791 20.274 16.84c12.633.048 17.726-8.31 17.726-8.31"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.049 17.658v12.684m9.619-12.554l-6.085 6.085l6.368 6.367"/><circle cx="42.1" cy="24" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}