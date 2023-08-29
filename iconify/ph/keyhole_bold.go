package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyholeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm0 192a84 84 0 1 1 84-84a84.09 84.09 0 0 1-84 84Zm0-144a44 44 0 0 0-32.94 73.16l-9.92 24.92A16 16 0 0 0 100 188h56a16 16 0 0 0 14.86-21.92l-9.92-24.92A44 44 0 0 0 128 68Zm7.59 74.38L144.2 164h-32.4l8.61-21.62a12 12 0 0 0-4.11-14.16a20 20 0 1 1 23.4 0a12 12 0 0 0-4.11 14.16Z"/>`),
		g.Group(children),
	)
}