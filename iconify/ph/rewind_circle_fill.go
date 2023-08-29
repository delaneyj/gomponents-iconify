package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.11 104.11 0 0 0 128 24Zm48 132a8 8 0 0 1-12.59 6.55l-40-28A8 8 0 0 1 120 128v28a8 8 0 0 1-12.59 6.55l-40-28a8 8 0 0 1 0-13.1l40-28A8 8 0 0 1 120 100v28a8 8 0 0 1 3.41-6.55l40-28A8 8 0 0 1 176 100Z"/>`),
		g.Group(children),
	)
}