package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.13 104.13 0 0 0 128 24Zm-16 136a8 8 0 0 1-16 0V96a8 8 0 0 1 16 0Zm48 0a8 8 0 0 1-16 0V96a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}