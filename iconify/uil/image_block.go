package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageBlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.54 2.46A5 5 0 1 0 22 6a5 5 0 0 0-1.46-3.54ZM14 6a3 3 0 0 1 3-3a3 3 0 0 1 1.29.3l-4 4A3 3 0 0 1 14 6Zm5.12 2.12a3.08 3.08 0 0 1-3.4.57l4-4A3 3 0 0 1 20 6a3 3 0 0 1-.88 2.12ZM19 13a1 1 0 0 0-1 1v.39l-1.48-1.49a2.87 2.87 0 0 0-3.93 0l-.7.71l-2.48-2.49a2.87 2.87 0 0 0-3.93 0L4 12.61V7a1 1 0 0 1 1-1h4a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 .95-.17h.09A3 3 0 0 0 20 19.44a1.43 1.43 0 0 0 0-.22V14a1 1 0 0 0-1-1ZM5 20a1 1 0 0 1-1-1v-3.57l2.9-2.89a.79.79 0 0 1 1.09 0l3.19 3.18L15.46 20Zm13-1a1 1 0 0 1-.18.54L13.3 15l.71-.7a.79.79 0 0 1 1.09 0l2.9 2.91Z"/>`),
		g.Group(children),
	)
}