package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlienFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 16a96.11 96.11 0 0 0-96 96c0 24 12.56 55.06 33.61 83c21.18 28.15 44.5 45 62.39 45s41.21-16.81 62.39-45c21.05-28 33.61-59 33.61-83a96.11 96.11 0 0 0-96-96ZM64 116a12 12 0 0 1 12-12a36 36 0 0 1 36 36a12 12 0 0 1-12 12a36 36 0 0 1-36-36Zm80 84h-32a8 8 0 0 1 0-16h32a8 8 0 0 1 0 16Zm12-48a12 12 0 0 1-12-12a36 36 0 0 1 36-36a12 12 0 0 1 12 12a36 36 0 0 1-36 36Z"/>`),
		g.Group(children),
	)
}