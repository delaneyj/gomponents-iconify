package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeControlTelephone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M15.553 10.972c2.08-2.194 2.08-5.75 0-7.944M17.95 13.5c3.402-3.59 3.402-9.41 0-13M13.5 8.806c.945-.998.945-2.614 0-3.612M6.94 23.5c-4.587-5.247-4.587-13.753 0-19h3.56v4.75H7.534v9.5H10.5v4.75H6.94Z"/>`),
		g.Group(children),
	)
}