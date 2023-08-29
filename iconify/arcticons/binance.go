package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Binance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 3.495l20.499 20.499l-20.5 20.499l-20.498-20.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.334 17.164l6.83 6.83l-6.83 6.832l-6.83-6.831zM24 17.17L30.83 24L24 30.832l-6.83-6.83z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.66 17.17L44.49 24l-6.83 6.83L30.828 24z"/>`),
		g.Group(children),
	)
}