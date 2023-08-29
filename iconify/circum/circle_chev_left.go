package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleChevLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.15 12.35a.492.492 0 0 1 0-.7l3-3a.495.495 0 0 1 .7.7L11.21 12l2.64 2.65a.495.495 0 0 1-.7.7Z"/><path fill="currentColor" d="M12 21.933A9.933 9.933 0 1 1 21.933 12A9.944 9.944 0 0 1 12 21.933Zm0-18.866A8.933 8.933 0 1 0 20.933 12A8.943 8.943 0 0 0 12 3.067Z"/>`),
		g.Group(children),
	)
}