package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayPauseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 64v128a8 8 0 0 1-16 0V64a8 8 0 0 1 16 0Zm40-8a8 8 0 0 0-8 8v128a8 8 0 0 0 16 0V64a8 8 0 0 0-8-8Zm-87.33 58.66L48.48 58.51A15.91 15.91 0 0 0 24 71.85v112.3A15.83 15.83 0 0 0 32.23 198a15.95 15.95 0 0 0 16.25-.53l88.19-56.15a15.8 15.8 0 0 0 0-26.68Z"/>`),
		g.Group(children),
	)
}