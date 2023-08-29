package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleChevUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.65 10.15a.492.492 0 0 1 .7 0l3 3a.495.495 0 0 1-.7.7L12 11.21l-2.65 2.64a.495.495 0 0 1-.7-.7Z"/><path fill="currentColor" d="M2.067 12A9.933 9.933 0 1 1 12 21.934A9.944 9.944 0 0 1 2.067 12Zm18.866 0A8.933 8.933 0 1 0 12 20.934A8.943 8.943 0 0 0 20.933 12Z"/>`),
		g.Group(children),
	)
}