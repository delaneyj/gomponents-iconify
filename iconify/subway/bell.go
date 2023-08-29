package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M419.7 340.3V169.6C419.7 75.4 343.3-1 249-1S78.3 75.4 78.3 169.7v170.7c0 42.7-42.7 85.3-42.7 85.3h426.7s-42.6-42.7-42.6-85.4zM249 511c35.4 0 64-19.1 64-42.7H185c0 23.6 28.6 42.7 64 42.7z"/>`),
		g.Group(children),
	)
}