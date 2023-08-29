package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualizerBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M80 108a12 12 0 0 1-12 12H24a12 12 0 0 1 0-24h44a12 12 0 0 1 12 12Zm-12 28H24a12 12 0 0 0 0 24h44a12 12 0 0 0 0-24Zm0 40H24a12 12 0 0 0 0 24h44a12 12 0 0 0 0-24Zm82-40h-44a12 12 0 0 0 0 24h44a12 12 0 0 0 0-24Zm0 40h-44a12 12 0 0 0 0 24h44a12 12 0 0 0 0-24Zm38-96h44a12 12 0 0 0 0-24h-44a12 12 0 0 0 0 24Zm44 16h-44a12 12 0 0 0 0 24h44a12 12 0 0 0 0-24Zm0 40h-44a12 12 0 0 0 0 24h44a12 12 0 0 0 0-24Zm0 40h-44a12 12 0 0 0 0 24h44a12 12 0 0 0 0-24Z"/>`),
		g.Group(children),
	)
}