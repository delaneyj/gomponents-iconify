package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jada(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.009 31.602a11.99 11.99 0 0 0 23.981 0V4.5H24.326v27.026"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.991 4.5l-11.664 7.573L35.99 18.13l-11.663 6.8l11.663 6.84l-11.663-.246l8.208 8.434m-8.209-8.434l-8.52 8.764"/>`),
		g.Group(children),
	)
}