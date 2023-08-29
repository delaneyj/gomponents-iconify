package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 78a30 30 0 1 0-30-30a30 30 0 0 0 30 30Zm0-48a18 18 0 1 1-18 18a18 18 0 0 1 18-18Zm101.14 101.09a6 6 0 0 1-8.23 2c-.37-.21-36.49-21.43-86.91-23v39.61l62.48 70.3a6 6 0 0 1-9 8L128 161l-59.52 67a6 6 0 0 1-9-8L122 149.72v-39.61c-50.42 1.6-86.55 22.82-86.92 23a6 6 0 0 1-6.17-10.29C30.6 121.84 71 98 128 98s97.39 23.84 99.09 24.86a6 6 0 0 1 2.05 8.23Z"/>`),
		g.Group(children),
	)
}