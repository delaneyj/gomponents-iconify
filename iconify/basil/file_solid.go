package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 2.25a.25.25 0 0 1 .25.25v5.647c0 .414.336.75.75.75h4.5a.25.25 0 0 1 .25.25V19A2.75 2.75 0 0 1 17 21.75H7A2.75 2.75 0 0 1 4.25 19V5A2.75 2.75 0 0 1 7 2.25h7Z"/><path fill="currentColor" d="M16.086 2.638c-.143-.115-.336.002-.336.186v4.323c0 .138.112.25.25.25h3.298c.118 0 .192-.124.124-.22L16.408 2.98a1.748 1.748 0 0 0-.322-.342Z"/>`),
		g.Group(children),
	)
}