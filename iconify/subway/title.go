package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Title(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M46.5 0v139.6h23.3c0-23.3 0-69.8 23.3-93.1c23.2-23.3 46.5-23.3 69.8-23.3h46.5v395.6c0 34.9-11.6 69.8-46.5 69.8h-22.8l-.5 23.2h232.7v-23.3H349c-34.9 0-46.5-34.9-46.5-69.8V23.3H349c23.3 0 46.5 0 69.8 23.3s23.3 69.8 23.3 93.1h23.3V0H46.5z"/>`),
		g.Group(children),
	)
}