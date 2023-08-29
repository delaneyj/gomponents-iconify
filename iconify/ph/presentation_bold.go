package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PresentationBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 36h-76V24a12 12 0 0 0-24 0v12H40a20 20 0 0 0-20 20v120a20 20 0 0 0 20 20h31l-16.4 20.5a12 12 0 0 0 18.74 15l28.4-35.5h52.46l28.4 35.5a12 12 0 0 0 18.74-15L185 196h31a20 20 0 0 0 20-20V56a20 20 0 0 0-20-20Zm-4 136H44V60h168Z"/>`),
		g.Group(children),
	)
}