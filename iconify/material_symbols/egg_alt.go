package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EggAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 22q-1.675 0-2.538-.563T12 20.25q-.475-.5-.913-.875T9.976 19q-1.125 0-2.5-.388t-2.587-1.274q-1.213-.888-2.038-2.313t-.85-3.5q-.05-4.175 2.063-6.85T9.974 2q1.775 0 3 .513T15.088 3.8q.887.775 1.5 1.713t1.162 1.812q.3.5.6.913T19 9q1.5 1.5 2.25 2.625t.75 3.4q0 3-1.863 4.988T16 22Zm-4-6.5q1.45 0 2.475-1.025T15.5 12q0-1.45-1.025-2.475T12 8.5q-1.45 0-2.475 1.025T8.5 12q0 1.45 1.025 2.475T12 15.5Z"/>`),
		g.Group(children),
	)
}