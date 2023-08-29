package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20ZM79.57 196.57a60 60 0 0 1 96.86 0a83.72 83.72 0 0 1-96.86 0ZM100 120a28 28 0 1 1 28 28a28 28 0 0 1-28-28Zm94 59.94a83.48 83.48 0 0 0-29-23.42a52 52 0 1 0-74 0a83.48 83.48 0 0 0-29 23.42a84 84 0 1 1 131.9 0Z"/>`),
		g.Group(children),
	)
}