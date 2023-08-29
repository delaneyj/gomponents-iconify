package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyDollarSimpleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm20 168h-12v8a8 8 0 0 1-16 0v-8h-8a36 36 0 0 1-36-36a8 8 0 0 1 16 0a20 20 0 0 0 20 20h36a20 20 0 0 0 0-40h-32a36 36 0 0 1 0-72h4v-8a8 8 0 0 1 16 0v8h4a36 36 0 0 1 36 36a8 8 0 0 1-16 0a20 20 0 0 0-20-20h-24a20 20 0 0 0 0 40h32a36 36 0 0 1 0 72Z"/>`),
		g.Group(children),
	)
}