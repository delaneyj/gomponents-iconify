package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OptionBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 184a12 12 0 0 1-12 12h-63.06a19.89 19.89 0 0 1-17.88-11.06L92.58 84H32a12 12 0 0 1 0-24h63.06a19.89 19.89 0 0 1 17.88 11.06L163.42 172H224a12 12 0 0 1 12 12ZM152 84h72a12 12 0 0 0 0-24h-72a12 12 0 0 0 0 24Z"/>`),
		g.Group(children),
	)
}