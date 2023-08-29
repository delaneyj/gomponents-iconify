package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodepenLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 13.202L13 15.535v3.596L19.197 15L16.5 13.202ZM14.697 12L12 10.202L9.303 12L12 13.798L14.697 12ZM20 10.868L18.303 12L20 13.132v-2.264ZM19.197 9L13 4.869v3.596l3.5 2.333L19.197 9ZM7.5 10.798L11 8.465V4.869L4.803 9L7.5 10.798ZM4.803 15L11 19.131v-3.596l-3.5-2.333L4.803 15ZM4 13.132L5.697 12L4 10.868v2.264ZM2 9a1 1 0 0 1 .445-.832l9-6a1 1 0 0 1 1.11 0l9 6A1 1 0 0 1 22 9v6a1 1 0 0 1-.445.832l-9 6a1 1 0 0 1-1.11 0l-9-6A1 1 0 0 1 2 15V9Z"/>`),
		g.Group(children),
	)
}