package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuotesBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M100 52H40a20 20 0 0 0-20 20v64a20 20 0 0 0 20 20h56v4a28 28 0 0 1-28 28a12 12 0 0 0 0 24a52.06 52.06 0 0 0 52-52V72a20 20 0 0 0-20-20Zm-4 80H44V76h52Zm120-80h-60a20 20 0 0 0-20 20v64a20 20 0 0 0 20 20h56v4a28 28 0 0 1-28 28a12 12 0 0 0 0 24a52.06 52.06 0 0 0 52-52V72a20 20 0 0 0-20-20Zm-4 80h-52V76h52Z"/>`),
		g.Group(children),
	)
}