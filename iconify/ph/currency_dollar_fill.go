package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyDollarFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M116 80h4v40h-4a20 20 0 0 1 0-40Zm32 56h-12v40h12a20 20 0 0 0 0-40Zm84-8A104 104 0 1 1 128 24a104.11 104.11 0 0 1 104 104Zm-48 28a36 36 0 0 0-36-36h-12V80h4a20 20 0 0 1 20 20a8 8 0 0 0 16 0a36 36 0 0 0-36-36h-4v-8a8 8 0 0 0-16 0v8h-4a36 36 0 0 0 0 72h4v40h-8a20 20 0 0 1-20-20a8 8 0 0 0-16 0a36 36 0 0 0 36 36h8v8a8 8 0 0 0 16 0v-8h12a36 36 0 0 0 36-36Z"/>`),
		g.Group(children),
	)
}