package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Walgreens(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 16.5c4.42-4.877 11.095.884 14.935-3.505c-4.053 5.883-15.522 19.377-7.223 22.434c9.268 3.414 24.536-15.911 19.263-19.142c-5.13-3.143-15.454 18.532-4.785 20.026c9.494 1.329 19.355-18.502 16.215-24.69"/>`),
		g.Group(children),
	)
}