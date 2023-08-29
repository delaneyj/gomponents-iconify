package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M160 180a36 36 0 1 0-36-36a36 36 0 0 0 36 36Zm0-48a12 12 0 1 1-12 12a12 12 0 0 1 12-12Zm56-64H113.76l81.69-24.5a12 12 0 0 0-6.9-23l-160 48A12 12 0 0 0 20 80v120a20 20 0 0 0 20 20h176a20 20 0 0 0 20-20V88a20 20 0 0 0-20-20Zm-4 128H44V92h168ZM60 124a12 12 0 0 1 12-12h24a12 12 0 0 1 0 24H72a12 12 0 0 1-12-12Zm0 40a12 12 0 0 1 12-12h24a12 12 0 0 1 0 24H72a12 12 0 0 1-12-12Z"/>`),
		g.Group(children),
	)
}