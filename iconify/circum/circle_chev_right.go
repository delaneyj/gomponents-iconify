package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleChevRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.85 11.65a.492.492 0 0 1 0 .7l-3 3a.495.495 0 0 1-.7-.7L12.79 12l-2.64-2.65a.495.495 0 0 1 .7-.7Z"/><path fill="currentColor" d="M12 2.067A9.933 9.933 0 1 1 2.067 12A9.944 9.944 0 0 1 12 2.067Zm0 18.866A8.933 8.933 0 1 0 3.067 12A8.943 8.943 0 0 0 12 20.933Z"/>`),
		g.Group(children),
	)
}