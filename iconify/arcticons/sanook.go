package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sanook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.474 27.077a6.682 6.682 0 0 0 5.834 3.424c2.755 0 5.12-1.667 6.177-4.087l-16.341-2.98C5.017 21.26 3.896 13.014 8.02 8.779c3.693-3.79 8.313-4.946 19.34.075v-2.82h14.906v14.889H29.984m-.456-.001a6.682 6.682 0 0 0-5.834-3.424a6.565 6.565 0 0 0-6.103 4.082l16.267 2.985c9.127 2.174 10.248 10.42 6.122 14.656c-3.693 3.79-8.313 4.946-19.34-.075v2.82H5.733V27.075h12.284"/>`),
		g.Group(children),
	)
}