package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainsSun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 10a4 4 0 1 0-4-4a4 4 0 0 0 4 4Zm0-6a2 2 0 1 1-2 2a2 2 0 0 1 2-2Zm-1.15 8.47a1 1 0 0 0-1.7 0l-1 1.63l-3.29-5.6a1 1 0 0 0-1.72 0l-7 12A1 1 0 0 0 3 22h18a1 1 0 0 0 .85-1.53ZM10.45 20H4.74L10 11l2.94 5l-1.25 2Zm2.35 0l1.49-2.37l.71-1.06l1-1.68L19.2 20Z"/>`),
		g.Group(children),
	)
}