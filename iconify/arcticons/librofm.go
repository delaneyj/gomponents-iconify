package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Librofm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.81 22.952l12.053-12.053v19.384L26.811 42.336a3.975 3.975 0 0 1-5.622 0L9.137 30.283V10.9l12.052 12.053a3.975 3.975 0 0 0 5.622 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.834 17.298a6.126 6.126 0 0 0-7.826-.02v.02m11.562-4.246a12.045 12.045 0 0 0-15.33 0m19.013-4.458a17.807 17.807 0 0 0-22.72 0"/>`),
		g.Group(children),
	)
}