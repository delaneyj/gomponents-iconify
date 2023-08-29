package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 256q0 40-17 75t-48 59l-20 122H85L65 390Q0 339 0 256t65-134L85 0h171l20 122q31 24 48 59t17 75zm-298 0q0 53 37.5 90.5T171 384t90.5-37.5T299 256t-37.5-90.5T171 128t-90.5 37.5T43 256z"/>`),
		g.Group(children),
	)
}