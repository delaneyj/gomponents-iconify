package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm0 192a88 88 0 1 1 88-88a88.1 88.1 0 0 1-88 88ZM112 96v64a8 8 0 0 1-16 0V96a8 8 0 0 1 16 0Zm48 0v64a8 8 0 0 1-16 0V96a8 8 0 0 1 16 0Z"/>`),
		g.Group(children),
	)
}