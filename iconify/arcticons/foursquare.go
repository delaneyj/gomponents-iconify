package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foursquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M10.01 43.07L9.705 5.841a1.056 1.056 0 0 1 1.046-1.065L37.42 4.5a.867.867 0 0 1 .86 1.025l-1.774 9.496a1.27 1.27 0 0 1-1.247 1.036H22.904a1.185 1.185 0 0 0-1.177 1.047l-.193 1.648a.938.938 0 0 0 .822 1.04a.948.948 0 0 0 .109.007h11.996a.846.846 0 0 1 .825 1.028l-2.02 9.088a1.314 1.314 0 0 1-1.282 1.029h-8.078a2.497 2.497 0 0 0-1.79.755l-11.362 11.67a.433.433 0 0 1-.744-.299Z"/>`),
		g.Group(children),
	)
}