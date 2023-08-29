package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 11.5a.5.5 0 0 1 0 1H9a.5.5 0 0 1 0-1Z"/><path fill="currentColor" d="M12 21.934A9.933 9.933 0 1 1 21.932 12A9.945 9.945 0 0 1 12 21.934Zm0-18.866A8.933 8.933 0 1 0 20.932 12A8.944 8.944 0 0 0 12 3.068Z"/>`),
		g.Group(children),
	)
}