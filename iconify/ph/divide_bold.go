package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DivideBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 128a12 12 0 0 1-12 12H40a12 12 0 0 1 0-24h176a12 12 0 0 1 12 12ZM128 84a20 20 0 1 0-20-20a20 20 0 0 0 20 20Zm0 88a20 20 0 1 0 20 20a20 20 0 0 0-20-20Z"/>`),
		g.Group(children),
	)
}